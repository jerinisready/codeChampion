// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package execution

import (
    // "bufio"
    "fmt"
    "os/exec"
    "time"

    // "strconv"
    // "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        fmt.Println(e.Error())
    }
}

func Complier(FileName string, Script string, Language string, Input string, Output string, Input2 string, Output2 string) string {
  t := time.Now()
  folder_name := "champ_" + (t.Format("20060102150405"))

    // To create directory
   os.Mkdir("/tmp/"+folder_name, 0777)
   f, err := os.Create("/tmp/"+folder_name + "/" + FileName)
   check(err)
   i, err3 := os.Create("/tmp/"+folder_name + "/in.txt")
   check(err3)

   // To write String
   _, err1 := f.WriteString(Script)
   check(err1)
   f.Sync()
   _, err2 := i.WriteString(Input)
   check(err2)
   i.Sync()
   defer f.Close()
   defer i.Close()
   path := "  /tmp/"+ folder_name + "/"
  //  var cmd1 string
  //  var cmd2 string
   if Language == "python" {
       cmd := "python "+ path + FileName + " < " + path + "in.txt "
       out_data, err := exec.Command(cmd).Output()
       if string(out_data) == "" {
            return err.Error()
         }
      if string(out_data) == Output{ return "" } else{ return string(out_data) }
    }
      if Language == "c" {
        cmd1 := "gcc -o hello " + path + FileName
        cmd2 := "./" + path + "hello " + " < " + path + "in.txt "
        out, _ := exec.Command(cmd1).Output()
        if out == nil{
          out, _ := exec.Command(cmd2).Output()
          if string(out) == Output{ return "" } else{ return  string(out)  }

        }
       if string(out) == Output{ return "" } else{ return  string(out)  }
     }
        if Language == "c++"{
         cmd1 := "g++ -o hello " + path + FileName
         cmd2 := "./" + path + "hello " + " < " + path + "in.txt "

         out, _ := exec.Command(cmd1).Output()
         if out == nil{
           out, _ := exec.Command(cmd2).Output()
           if string(out) == Output{ return "" } else{ return  string(out)  }

         }
        if  string(out)  == Output{ return "" } else{ return  string(out)  }
      }
      // else if Language == "swift"{
      //     cmd1 = "swiftc " + path + FileName
      //     cmd2 = "./" + path + FileName[:6] +  " < " + path + "in.txt "
      //       out, _ := exec.Command(cmd1).Output()
      //     if out == nil{
      //       out, _ := exec.Command(cmd2).Output()
      //     }
      //    if out == Output{ return "" } else{ return out }
      //  }
        if Language == "java"{
          cmd1 := "javac " + path + FileName
          cmd2 := "java " + path + FileName[:5] + " < " + path + "in.txt "

          out, _ := exec.Command(cmd1).Output()
          if  out  == nil{
            out, _ := exec.Command(cmd2).Output()
            if string(out) == Output{ return "" } else{ return  string(out)  }

          }
         if  string(out)  == Output{ return "" } else{ return  string(out)  }
       }
       if Language == "ruby"{
           cmd2 := "ruby " + path + FileName[:5] + " < " + path + "in.txt "
           out, _ := exec.Command(cmd2).Output()
           if  string(out)  == Output{ return "" } else{ return  string(out)  }
        }
       if Language == "go"{
            cmd2 := "go run " + path + FileName[:5] + " < " + path + "in.txt "
            out, _ := exec.Command(cmd2).Output()
            if  string(out)  == Output{ return "" } else{ return  string(out)  }
         }
         if Language == "elixir"{
             cmd2 := "elixirc " + path + FileName[:5] + " < " + path + "in.txt "
             out, _ := exec.Command(cmd2).Output()
             if  string(out)  == Output{ return "" } else{ return  string(out)  }
          }
         if Language == "lisp"{
              cmd2 := "clisp " + path + FileName[:5] + " < " + path + "in.txt "
              out, _ := exec.Command(cmd2).Output()
              if  string(out)  == Output{ return "" } else{ return  string(out)  }
           }
         if Language == "cobol"{
           cmd1 := "cobc -free -x -o executable-program " + path + FileName[:5] + " < " + path + "in.txt "
           cmd2 := "./" + path + "executable-program " + " < " + path + "in.txt "
           out, _ := exec.Command(cmd1).Output()
           if  out == nil{
             out, _ := exec.Command(cmd2).Output()
             if string(out) == Output{ return "" } else{ return  string(out)  }

           }
           if  string(out)  == Output{ return "" } else{ return  string(out)  }
          }
          return "#"
}
//ELIXIR : elixirc hello.ex
//COBOL :  cobc -free -x -o helloworld-exe helloworld
//LISP : clisp hello.lisp
// main("ss", "ssssss", "ph", "11", "33")
