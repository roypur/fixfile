package main

import (
  "bufio"
  "path/filepath"
  "io/ioutil"
  "os"
  "os/exec"
  "bytes"
  "strings"
)

func main(){
    
    const space string = string(32)
    
    var files []string = make([]string, 0)
    
    for k,v := range os.Args{
        if k > 0{
            file, err := filepath.Abs(v)
        
            if (err == nil) && (isText(file)) {
                files = append(files, file)
            }
        }
    }

    for _,fileName := range files{
    

        file,_ := os.Open(fileName)
    
        defer file.Close()

        var txt string = ""
  
        scanner := bufio.NewScanner(file)
    
        for scanner.Scan() {
            var next string = scanner.Text()
        
            var nextLine string = ""

            // cleaning the file
            for _,v := range next{
        
                // remove garbage
                if( ( (v < 846) && (v > 160) && (v != 173) ) || ( (v < 127) && (v > 32) ) || (v == 10) || (v == 9) || ( (v < 9210) && (v > 8592) ) || ( (v < 11193) && (v > 9312) ) ){
                    nextLine += string(v)
                }else{
                    nextLine += space
                }
            
            }


            var trimmed string = strings.TrimRight(nextLine, space)
            
            if len(trimmed) == 0{
                trimmed = nextLine
                // dont trim if this is a line with only indentation
            }

            txt += trimmed
            txt += "\n"
        }
        
        file.Close()

        ioutil.WriteFile(fileName, []byte(txt), os.ModePerm)
    }
    
}

func isText(filename string)(bool){

    var out bytes.Buffer

    cmd := exec.Command("/usr/bin/file", "--mime-type", filename)

    cmd.Stdout = &out

    err := cmd.Run()
    
    if err != nil{
        return false
    }

    var mime []string = strings.Split(out.String(), ":")
    
    var fileType string = strings.TrimSpace( strings.Split(mime[1], "/")[0] ) 

    if fileType == "text"{
        return true
    }

    return false

}

