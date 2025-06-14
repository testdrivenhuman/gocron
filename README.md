
# ⏰ gocron

A lightweight, zero-dependency, custom-built cron job scheduler in Go — inspired by `node-schedule`

🚀 Supports full cron expressions like `*/5 * * * * *` and schedules multiple async jobs with efficiency.

---

### 📦 Installation

```bash
go get github.com/testdrivenhuman/gocron
```




### 📄 Import

```
import "github.com/testdrivenhuman/gocron/cron"
```

### 🛠️ Usage

```
package main

import (
    "fmt"
    "github.com/testdrivenhuman/gocron/cron"
)

func main() {
    scheduler := cron.NewScheduler()

    scheduler.AddJob("*/5 * * * * *", func() {
        fmt.Println("Job runs every 5 seconds!")
    })

    scheduler.Start()

    select {} // Keeps the program running
}

```

### 🧠 Cron Expression Format
```

┌────────────── second (0-59)
│ ┌──────────── minute (0 - 59)
│ │ ┌────────── hour (0 - 23)
│ │ │ ┌──────── day of month (1 - 31)
│ │ │ │ ┌────── month (1 - 12)
│ │ │ │ │ ┌──── day of week (0 - 6) (Sunday=0)
│ │ │ │ │ │
│ │ │ │ │ │
* * * * * *

Examples

Expression	Description
*/2 * * * * *	Every 2 seconds
0 */1 * * * *	Every 1 minute
0 0 * * * *	Every hour
0 30 10 * * *	10:30:00 every day

```

## ✨ Features

- ✅ Native cron string parser  
- ✅ Schedule multiple jobs  
- ✅ Async execution (non-blocking)  
- ✅ Built from scratch (no third-party dependencies)  
- ✅ CPU efficient (single ticker loop)  
- 🛠️ Upcoming Features:
  - Cancel jobs  
  - Job logging  
  - Retry mechanism  
  - Error handling  

### 📁 Project Structure
```
gocron/
├── cron/
│   └── scheduler.go    # Core scheduler logic
├── main.go             # Entry point (example/test)
└── go.mod
```

<br>

### 📥 Contributing

Feel free to contribute by submitting issues, feature requests, or PRs!

To contribute:

#### Fork and clone the repo
```git clone https://github.com/testdrivenhuman/gocron.git```

#### Create a branch
```git checkout -b feature/some-feature```

#### Commit your changes
```git commit -am "Added awesome feature"```

#### Push and create a PR!
```git push origin feature/some-feature```

<br>
<br>

## 📃 License

MIT License © testdrivenhuman
<br>
<br>

### ❤️ Show Some Love

If you like this project:
	•	🌟 Star it on GitHub
	•	🧵 Share it on Twitter/LinkedIn
	•	🛠️ Use it in your projects

---
