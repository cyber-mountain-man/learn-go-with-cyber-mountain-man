package config

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"  // Gorilla's secure session library
	"github.com/joho/godotenv"     // .env loader for local development
)

// Store holds the global session store instance
var Store *sessions.CookieStore

func init() {
	// ✅ Load .env in development (optional in production environments)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env not found — using system environment variables instead")
	}

	// 🔐 Retrieve session keys from environment
	authKey := os.Getenv("SESSION_AUTH_KEY")
	encryptKey := os.Getenv("SESSION_ENCRYPT_KEY")

	// ❌ Abort if keys are missing (required for session security)
	if authKey == "" || encryptKey == "" {
		log.Fatal("❌ SESSION_AUTH_KEY and SESSION_ENCRYPT_KEY must be set")
	}

	// ✅ Initialize the Gorilla session store with secure keys
	Store = sessions.NewCookieStore([]byte(authKey), []byte(encryptKey))
	Store.Options = &sessions.Options{
		Path:     "/",                     // Session cookie is valid site-wide
		MaxAge:   3600,                    // 1-hour expiration
		HttpOnly: true,                    // JS can't access
		SameSite: http.SameSiteStrictMode, // CSRF protection
		// Secure: true,                   // Uncomment when deploying over HTTPS
	}
}

/*
🔐 SESSION SECURITY CONFIGURATION (Gorilla + .env)

✅ What Happens Here:
- This file defines a global, shared Gorilla `CookieStore` using **secure auth/encryption keys**.
- Session keys are loaded from the environment — via `.env` for dev or OS variables in production.
- The store is configured with strict cookie options for session security.

✅ Why This Matters:
- Session cookies contain sensitive data (e.g., login state), so integrity and confidentiality are crucial.
- This setup ensures **tamper-proof** and **encrypted** session data.
- Externalized secrets support scalable, secure deployments across dev, staging, and prod.

✅ Key Concepts:
| Concept                      | Purpose                                          |
|------------------------------|--------------------------------------------------|
| `sessions.NewCookieStore()`  | Creates a tamper-resistant, encrypted session store |
| `.env` + `os.Getenv()`       | Loads session keys securely without hardcoding   |
| `HttpOnly`, `SameSite`       | Mitigate XSS and CSRF attacks                   |

🚨 Security Tip:
- Do **not** commit `.env` files or keys to version control.
- Use `openssl rand -hex 32` to generate strong keys.
- For production, use a secrets manager (e.g., AWS Secrets Manager, Vault, GCP Secret Manager).

📚 Up Next:
- Use `config.Store.Get()` in your handlers and middleware to manage sessions securely.
- Implement flash messages or user roles with `session.Values[...]`.
*/
