
package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

//  The spec says Python standard library plus argparse only, but can non-ASCII bytes be counted without them?

//  Go standard library is lighter due to static compilation; have to explicitly import more.

//  Why is the world still using wc written in C?
//    Backwards compat - see your point.


func main() {

    var args []bool

    var filename string

    var data []byte

    var chars, words, lines, bytec int

    fmt.Println(os.Args[1:])

    args, filename = handle_args(os.Args[1:])

    data = load_data(filename)

    chars, words, lines, bytec = count_chars(data), count_words(data), count_lines(data), len(data)

    output_data(chars, words, lines, bytec, args)

}


func handle_args(all_args []string) {

    join_args := strings.Join(all_args, " ")

    fmt.Println(join_args)

    split_args := strings.Split(join_args, " ")

    //  Assume last argument is filename.

    filename := split_args[len(split_args)-1]
    
    var m_flag, w_flag, l_flag, c_flag, big_l_flag bool = false, false, false, false, false

    for _, char := range split_args[0] {

        switch char {

        case 'm':
            m_flag = true

        case 'w':
            w_flag = true

        case 'l':
            l_flag = true

        case 'c':
            c_flag = true

        case 'L':
            big_l_flag = true

        }
    }

    //  If all false, set flags to default context.

    if (!m_flag && !w_flag && !l_flag && !c_flag && !big_l_flag) {

        m_flag, w_flag, l_flag, c_flag = true, true, true, true
    
    }

    return []bool{m_flag, w_flag, l_flag, c_flag, big_l_flag}, filename

}


func load_data(filename string) {

    data, _ := ioutil.ReadFile(filename)

    return data

}


func count_chars(data []byte) {

    //  Loop over text, count chars.

    var int ascii_count

    for _, b := range data {
        
        if b >= 32 && b <= 126 {
            
            ascii_count++
        
        }
    
    }

    return ascii_count

}


func count_words(data []byte) {

    var bool prev_byte_alpha

    var int word_count

    for _, b := range data {

        if (prev_byte_alpha == true && b == 10) {

            word_count++
        
        }
        
        if ( b >= 65 && b <= 90 ) || ( b >= 97 && b <= 122 ) {
            
            prev_byte_alpha = true
        
        } else {

            prev_byte_alpha = false
        
        }
    
    }

    return word_count

}


func count_lines(data []byte) {

    var int newlines_count

    for _, b := range data {
        
        if b == 10 {
            
            newlines_count++
        
        }
    
    }

    return newlines_count

}


func output_data(chars int, words int, lines int, bytec int, args []bool) {

          

}



