package runner

const (
  version = "1.0.0"
  author  = "Zeckers75"
  banner  = `

     /__/\          ___         ___        /  /\      ___          /__/|          ___
     \  \:\        /  /\       /  /\      /  /::\    /  /\        |  |:|         /__/|
      \__\:\      /  /:/      /  /:/     /  /:/\:\  /  /:/        |  |:|        |  |:|
  ___ /  /::\    /  /:/      /  /:/     /  /:/~/:/ /__/::\      __|__|:|        |  |:|
 /__/\  /:/\:\  /  /::\     /  /::\    /__/:/ /:/  \__\/\:\__  /__/::::\____  __|__|:|
 \  \:\/:/__\/ /__/:/\:\   /__/:/\:\   \  \:\/:/      \  \:\/\    ~\~~\::::/ /__/::::\
  \  \::/      \__\/  \:\  \__\/  \:\   \  \::/        \__\::/     |~~|:|~~     ~\~~\:\
   \  \:\           \  \:\      \  \:\   \  \:\        /__/:/      |  |:|         \  \:\
    \  \:\           \__\/       \__\/    \  \:\       \__\/       |  |:|          \__\/
     \__\/                                 \__\/                   |__|/
                                                                                       

      v` + version + ` - @` + author
  usage = `
  [buffers] | httpixy [options]
  httpixy [options]`
  options =  `
  -u, --url <URL>           Define single URL to test
  -l, --list <FILE>         Test URLs within file
  -X, --method <METHOD>     Specify request method to use (default: GET)
  -o, --output <FILE>       File to save results
  -x, --proxy <URL>         Use proxy
  -c, --concurrent <i>      Set the concurrency level (default: 20)
  -s, --silent              Silent mode
  -V, --version             Show current HTTPIXY version
  -h, --help                Display its help
`
)

/*
 -H, --header <HEADER>     Pass custom header to target
 -v, --verbose             Verbose mode
 -d, --data <DATA>         Define request data
 */