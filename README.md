# Terminal Portfolio

An interactive SSH-based terminal portfolio built with Go, featuring a beautiful TUI (Terminal User Interface) powered by Charm's Bubbletea and styled with Lipgloss.

## Features

- **SSH Server Integration** - Access portfolio via SSH from anywhere
- **Beautiful TUI** - Styled with Lipgloss for modern terminal aesthetics  
- **Horizontal Navigation** - Intuitive left/right arrow key navigation
- **Orange Theme** - Eye-catching orange accents for selected items
- **Responsive Design** - Adapts to different terminal sizes
- **Real-time Updates** - Smooth navigation with instant content switching
- **Multiple Sections** - Home, About, Projects, and Contact pages

## 🛠️ Tech Stack

- **[Go](https://golang.org/)** - Backend language
- **[Wish](https://github.com/charmbracelet/wish)** - SSH server framework
- **[Bubbletea](https://github.com/charmbracelet/bubbletea)** - TUI framework
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)** - Terminal styling

## 📋 Prerequisites

- Go 1.24+ installed
- Terminal with SSH support
- Basic understanding of Go modules

## 🚀 Installation & Setup

### 1. Clone the Repository
```bash
git clone https://github.com/Rudra-Sankha-Sinhamahapatra/Terminal-Portfolio
cd Terminal-Portfolio
```

### 2. Install Dependencies
```bash
go mod download
go mod tidy
```

### 3. Generate SSH Host Key
```bash
mkdir -p .ssh
ssh-keygen -t ed25519 -f .ssh/id_ed25519 -N ""
```

### 4. Clean Old SSH Keys (If Needed)
If you've run this before and getting "Host key verification failed":

```bash
# Remove old host key for localhost:23234
ssh-keygen -R "[localhost]:23234"

# Or remove all localhost entries
ssh-keygen -R "localhost"

```

### 5. Build and Run

#### Option 1: Run Directly (Development)
```bash
go run main.go
```

#### Option 2: Build Binary (Production)
```bash
go build -o portfolio main.go

./portfolio
```


#### Option 4: Install Globally
```bash
go install .

terminal-portfolio
```

## 🎮 Usage

### Starting the Server
```bash
go run main.go
```
You should see:
```
INFO Starting SSH server host=localhost port=23234
```

### Connecting via SSH
**From another terminal:**
```bash
ssh localhost -p 23234
```

**From external machine:**
```bash
ssh <your-ip> -p 23234
```

### Navigation Controls
- **←/→ Arrow Keys** - Navigate between menu items
- **q or Ctrl+C** - Quit application  
- **No Enter needed** - Content updates automatically

## 📁 Project Structure

```
terminal-portfolio/
├── main.go              # SSH server and application entry point
├── banner/
│   └── banner.go        # ASCII art banner
├── menu/
│   ├── menu.go          # Main TUI logic and navigation
│   ├── pages/           # Individual page implementations
│   │   ├── page.go      # Page interface definition
│   │   ├── home.go      # Home page content 
│   │   ├── about.go     # About page with skills
│   │   ├── projects.go  # Projects showcase
│   │   └── contact.go   # Contact information
│   ├── styles/          # Centralized styling
│   │   └── styles.go    # All styling definitions
│   └── utils/           # Utility functions
│       └── layout.go    # Layout helper functions
├── .ssh/
│   ├── id_ed25519       # SSH host private key  
│   └── id_ed25519.pub   # SSH host public key
├── go.mod               # Go module dependencies
├── go.sum               # Dependency checksums
├── .gitignore           # Git ignore rules
└── README.md            # This file
```

## 🔧 Customization

### Update Personal Information

#### Title and Navigation
- Edit `menu/menu.go` line 167: Change `"Rudra | Portfolio"`

#### Page Content
- **Home Page**: Edit `menu/pages/home.go` - Update banner subtitle
- **About Section**: Edit `menu/pages/about.go` - Update bio and skills arrays
- **Projects**: Edit `menu/pages/projects.go` - Modify the `getProjects()` method
- **Contact**: Edit `menu/pages/contact.go` - Update the `getContactInfo()` method

#### Banner
- **ASCII Art**: Edit `banner/banner.go` - Replace with your custom ASCII art


## 🐛 Troubleshooting

### "Host key verification failed"
```bash
ssh-keygen -R "[localhost]:23234"
ssh localhost -p 23234
```

### "Connection refused"
- Ensure the server is running (`go run main.go`)
- Check if port 23234 is available
- Try a different port in `main.go`

### "Permission denied"
- Make sure SSH keys have correct permissions:
```bash
chmod 600 .ssh/id_ed25519
chmod 644 .ssh/id_ed25519.pub
```

### "Module not found"
```bash
go mod download
go mod tidy
```


## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 🌟 Acknowledgments

- [Charm](https://charm.sh/) for the amazing TUI libraries
- [Bubbletea](https://github.com/charmbracelet/bubbletea) community
- [Lipgloss](https://github.com/charmbracelet/lipgloss) for beautiful styling

---

**Made with ❤️ by [Rudra Sankha Sinhamahapatra](https://github.com/Rudra-Sankha-Sinhamahapatra)** 