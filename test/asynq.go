package test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"
)

// Struct untuk Asynq
type Asynq struct {
	RedisOpt asynq.RedisClientOpt
	Client   *asynq.Client
}

func NewAsynq() *Asynq {
	return &Asynq{
		RedisOpt: asynq.RedisClientOpt{Addr: "localhost:6379"},
		Client:   asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"}),
	}
}

// Handler untuk job email:send
func handleEmailTask(ctx context.Context, t *asynq.Task) error {
	email := string(t.Payload()) // Ambil email dari payload
	fmt.Printf("📨 Mengirim email ke: %s...\n", email)

	// Simulasi error (50% kemungkinan gagal)
	if time.Now().Unix()%2 == 0 {
		return fmt.Errorf("❌ Gagal mengirim email ke %s", email)
	}

	fmt.Println("✅ Email berhasil dikirim!")
	return nil
}

// Handler untuk memproses task
func handleEmailTaskDelayedJob(ctx context.Context, t *asynq.Task) error {
	email := string(t.Payload()) // Ambil email dari payload
	fmt.Printf("📩 Mengirim email ke: %s...\n", email)
	return nil
}

// Worker (Harus dijalankan dalam proses terpisah)
func (a *Asynq) Worker() {
	// Konfigurasi Worker
	server := asynq.NewServer(a.RedisOpt, asynq.Config{
		Concurrency: 10, // Bisa menjalankan 10 job sekaligus
		Queues: map[string]int{
			"default": 1, // Prioritas queue "default"
		},
	})

	// Router untuk task yang ditangani Worker
	mux := asynq.NewServeMux()
	mux.HandleFunc("email:send", handleEmailTask)
	mux.HandleFunc("emailDelayed:send", handleEmailTaskDelayedJob)

	fmt.Println("🚀 Worker berjalan...")
	if err := server.Run(mux); err != nil {
		log.Fatal(err)
	}
}

// Client untuk menambahkan job ke antrian
func (a *Asynq) Send() {
	defer a.Client.Close()

	// Buat task untuk mengirim email
	email := "user@example.com"
	task := asynq.NewTask("email:send", []byte(email))

	// Tambahkan ke antrian dengan opsi retry & delay
	info, err := a.Client.Enqueue(task,
		asynq.MaxRetry(5),              // Retry maksimal 5 kali
		asynq.ProcessIn(5*time.Second), // Delay 5 detik sebelum dieksekusi
		asynq.Queue("default"),         // Masukkan ke queue "default"
	)

	if err != nil {
		log.Fatalf("❌ Gagal menambahkan job: %v", err)
	}

	log.Printf("✅ Job berhasil ditambahkan: ID=%s", info.ID)
}

func (a *Asynq) SendDelayedJob() {
	defer a.Client.Close()

	// Buat task baru untuk kirim email
	email := "user@example.com"
	task := asynq.NewTask("email:send", []byte(email))

	// Enqueue task dengan delay 1 menit
	info, err := a.Client.Enqueue(task, asynq.ProcessIn(1*time.Minute))
	if err != nil {
		log.Fatalf("❌ Gagal enqueue task: %v", err)
	}

	fmt.Printf("✅ Task akan dieksekusi dalam 1 menit. ID=%s\n", info.ID)
}

// Implement to cobra
var (
	asynqWorker bool
	asynqSend   bool
)

// command
var asynqCmd = &cobra.Command{
	Use:   "asynq",
	Short: "asynq",
	Long:  `asynq for using cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		default:
			cmd.Help()
		case asynqWorker:
			a := NewAsynq()
			a.Worker()
		case asynqSend:
			a := NewAsynq()
			a.Send()
		}
	},
}
