package main
import (
	"fmt"
	"io"
	"flag"
	"os"
	"bufio"
	"errors"
)
/*================================= types =========================*/
type selpg_args struct {
	startPage int
	endPage int
	page_len int
	auto_line bool
	print_dest string
}

/*================================= globals =======================*/

var progname string



/*================================= process_args() ================*/
func process_args(args *selpg_args) {
	if args.startPage == -1 || args.endPage == -1 {
		fmt.Fprintf(os.Stderr, "%s: 参数输入不足\n", progname)
		flag.Usage()
		os.Exit(1)
	}
	if args.startPage > args.endPage || args.startPage < 0 || args.endPage < 0 {
		fmt.Fprintf(os.Stderr, "%s: 页码范围错误\n", progname)
		flag.Usage()
		os.Exit(2)
	}
	if os.Args[1][0] != '-' || os.Args[1][1] != 's' {
		fmt.Fprintf(os.Stderr, "%s: 指令输入错误\n")
		flag.Usage()
		os.Exit(3)
	}
	end := 2
	if len(os.Args[1] == 2) {
		end = 3
	}

	if os.Args[end][0] != '-' || os.Args[end][1] != 'e' {
		fmt.Fprintf(os.Stderr, "%s: 参数输入不足或错误", progname)
		flag.Usage()
		os.Exit(4)
	}

	if args.auto_line == false && args.page_len == -1 {
		args.number = 72
	}

	if args.auto_line == true && args.number != -1 {
		fmt.Fprintf(os.Stderr, "%s: 换页方式冲突", progname)
		flag.Usage()
		os.Exit(5)
	}
}

/*================================= process_input() ===============*/
func process_input(args *selpg_args, in string, out string) {
	if in != "" {
		infile, err := os.Open(in)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		reader := bufio.NewReader(infile)
	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	if out != "" {
		writter, err := os.OpenFile(out, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		writter = nil
	}

	read_write(reader, writter)
}


func read_write(in *bufio.Reader, out *os.File) {
	if !args.auto_line {
		count := 0
		for true {
			line, _, err := in.ReadLine()
			if err != io.EOF && err != nil {
			fmt.Println(err)
			os.Exit(1)
				}
			if err == io.EOF {
				break
				}
			if count / args.page_len >= args.startPage {
				if count / args.page_len >= args.endPage {
					break
				} else {
				out.WriteString(line)
			}
		count++
		}
	} else {
		for pageNum := 0; pageNum <= args.endPage; pageNum++ {
		line, err := in.ReadString('\f')
		if err != io.EOF && err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err == io.EOF {
			break
		}
		out.WriteString(line)
	}
	}
}

/*================================= usage() =======================*/
func Usage() {
	fmt.Println("\nUSAGE: %s -sstart_page -eend_page [ -f | -llines_per_page ]"
	" [ -ddest ] [ in_filename ]\n", progname)
}

/*================================= main()=== =====================*/

func main() {
	progname = os.Args[0]
	var args selpg_args
	flagSet = flag.NewFlagSet(os.Args[0], flag.PanicOnError)
	flag.Usage = Usage
	flag.IntVar(&args.startPage, "s", -1, "start page")
	flag.IntVar(&args.endPage, "e", -1, "end page")
	flag.IntVar(&args.page_len, "l", -1, "the number of line")
	flag.BoolVar(&args.auto_line, "f", false, "wrap line automatically")
	flag.StringVar(&args.dest, "d", "", "to printer")
	flag.Parse()
	process_args(&args)
	var in string
	var out string
	if flag.Set.NArg() > 0 {
		in = flagSet.Arg(0)
	}
	if flagSet.NArg() > 1 {
		out = flagSet.Arg(1)
	}
	process_input(&args, in, out)
}