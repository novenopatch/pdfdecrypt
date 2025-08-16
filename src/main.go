package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

const version = "1.0.0"

//go:embed binaries/qpdf-darwin-arm64
var qpdfDarwinARM []byte

//go:embed binaries/qpdf-linux-amd64
var qpdfLinuxAMD []byte

//go:embed binaries/qpdf-windows-amd64.exe
var qpdfWindowsAMD []byte

func main() {
	helpFlag := flag.Bool("help", false, "Afficher l'aide")
	versionFlag := flag.Bool("version", false, "Afficher la version")
	passwordFlag := flag.String("password", "", "Mot de passe PDF")
	inputFlag := flag.String("file", "", "Fichier PDF spécifique à déchiffrer")
	dryRunFlag := flag.Bool("dry-run", false, "Lister les fichiers sans les déchiffrer")
	flag.Parse()

	if *helpFlag {
		fmt.Println("Usage: pdfdecrypt [options]")
		fmt.Println("Options:")
		fmt.Println("  --help            Afficher cette aide")
		fmt.Println("  --version         Afficher la version")
		fmt.Println("  --password=PASS   Mot de passe PDF (ou PDF_PASSWORD dans .env)")
		fmt.Println("  --file=FILE       Fichier PDF spécifique à déchiffrer")
		fmt.Println("  --dry-run         Lister les fichiers sans les déchiffrer")
		os.Exit(0)
	}

	if *versionFlag {
		fmt.Println("pdfdecrypt version", version)
		os.Exit(0)
	}
	err := godotenv.Load()
	if err != nil {
		log.Println("Pas de fichier .env trouvé, utilisation des valeurs par défaut ou arguments")
	}

	pdfPassword := *passwordFlag
	if pdfPassword == "" {
		pdfPassword = os.Getenv("PDF_PASSWORD")
	}
	if pdfPassword == "" {
		log.Fatal("⚠️  Mot de passe non fourni (via --password ou PDF_PASSWORD)")
	}
	srcDir := os.Getenv("PDF_SRC_DIR")
	if srcDir == "" {
		srcDir = "."
	}

	outDir := os.Getenv("PDF_OUT_DIR")
	if outDir == "" {
		outDir = "dercipts"
	}

	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		os.MkdirAll(outDir, os.ModePerm)
	}

	logFile, err := os.OpenFile("decrypt.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Impossible d'ouvrir decrypt.log: %v", err)
	}
	defer logFile.Close()
	logger := log.New(logFile, "", log.LstdFlags)

	qpdfPath := extractQPDF()
	var files []string
	if *inputFlag != "" {
		files = []string{*inputFlag}
	} else {
		files, err = filepath.Glob(filepath.Join(srcDir, "*.pdf"))
		if err != nil {
			log.Fatal(err)
		}
	}
	if *dryRunFlag {
		fmt.Println("📋 Dry-run mode: files that would be decrypted:")
		for _, f := range files {
			fmt.Println(" -", f)
		}
		return
	}
	for _, f := range files {
		decryptPDF(qpdfPath, pdfPassword, f, outDir, logger)
	}
}
func extractQPDF() string {
	var data []byte
	var name string

	switch runtime.GOOS {
	case "darwin":
		if runtime.GOARCH == "arm64" {
			data = qpdfDarwinARM
			name = "qpdf"
		}
	case "linux":
		if runtime.GOARCH == "amd64" {
			data = qpdfLinuxAMD
			name = "qpdf"
		}
	case "windows":
		if runtime.GOARCH == "amd64" {
			data = qpdfWindowsAMD
			name = "qpdf.exe"
		}
	default:
		log.Fatalf("❌ Platforme non supportée: %s/%s", runtime.GOOS, runtime.GOARCH)
	}

	tmpPath := filepath.Join(os.TempDir(), name)
	err := os.WriteFile(tmpPath, data, 0755)
	if err != nil {
		log.Fatalf("Impossible d'écrire le binaire qpdf: %v", err)
	}
	return tmpPath
}
func decryptPDF(qpdfPath, password, inputFile, outDir string, logger *log.Logger) {
	outFile := filepath.Join(outDir, filepath.Base(inputFile))
	cmd := exec.Command(qpdfPath, fmt.Sprintf("--password=%s", password), "--decrypt", inputFile, outFile)

	start := time.Now()
	fmt.Printf("Décryptage de %s → %s\n", inputFile, outFile)
	err := cmd.Run()
	duration := time.Since(start)
	if err != nil {
		msg := fmt.Sprintf("ERREUR: %s → %s : %v (durée: %s)", inputFile, outFile, err, duration)
		fmt.Println(msg)
		logger.Println(msg)
	} else {
		msg := fmt.Sprintf("SUCCES: %s → %s (durée: %s)", inputFile, outFile, duration)
		logger.Println(msg)
	}
}
