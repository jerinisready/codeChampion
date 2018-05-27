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

func Complier(FileName string, Script string, Language string, Input string, Output string) string {
  t := time.Now()
  folder_name := "champ_" + (t.Format("20060102150405"))

    // To create directory
   os.Mkdir("/tmp/"+folder_name, 0777)
   f, err := os.Create("/tmp/"+folder_name + "/" + FileName)
   check(err)
   // To create file
   i, err3 := os.Create("/tmp/"+folder_name + "/in.txt")
   check(err3)
   // _, err4 := os.Create("/tmp/"+folder_name + "/out.txt")
   // check(err4)

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

   if Language == "python"{
       cmd := "python "+ path + FileName + " < " + path + "in.txt "
       out_data, err := exec.Command(cmd).Output()
       check(err)
      if string(out_data) == Output{ return "" } else{ return string(out_data) }
    }
    // if Language == "c"{
    //     cmd1 = "gcc -o hello " + path + FileName
    //     cmd2 = "./" + path + "hello " + " < " + path + "in.txt > " + path + "out.txt"
    //
    //     out, err := exec.Command(cmd1).Output()
    //     out, err := exec.Command(cmd2).Output()
    //    if out == Output{ return "" } else{ return out }
    //  }
    //  if Language == "c++"{
    //      cmd1 = "g++ -o hello " + path + FileName
    //      cmd2 = "./" + path + "hello " + " < " + path + "in.txt > " + path + "out.txt"
    //
    //      out, err := exec.Command(cmd1).Output()
    //      out, err := exec.Command(cmd2).Output()
    //     if out == Output{ return "" } else{ return out }
    //   }
    //   if Language == "java"{
    //       cmd1 = "g++ -o hello " + path + FileName
    //       cmd2 = "./" + path + "hello " + " < " + path + "in.txt > " + path + "out.txt"
    //
    //       out, err := exec.Command(cmd1).Output()
    //       out, err := exec.Command(cmd2).Output()
    //      if out == Output{ return "" } else{ return out }
    //    }


  //
//   if(Language == "python"){
//     return true
//  fmt.Printf(folder_name, "oooooooooo")

//   }
//   else{
//     return false
//   }
return ""
}

// main("ss", "ssssss", "ph", "11", "33")
