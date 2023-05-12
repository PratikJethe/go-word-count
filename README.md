# Go Word Count
wc is an implementation of unix wc in golang

## Features

- word count, line count, character count and total count
- can walk through directory with multiple files
## Flags

wc supports various flags. multiple flags can be combined to generate desired output. 
If none of the flags are provided all flags are set to true by default

| flag | desrcription 
| ------ | ------ 
| -l | provides line count 
| -w |  provides word count
| -c |  provides character count


## How to use
- clone this repository [git clone https://github.com/PratikJethe/go-word-count]
- run [go build -o wc.exe] in root of project
- execute binary with appropriate input

## Examples 

### 1. single text file
 
```sh
$ ./wc.exe -c -l -w  test-data/Hamlet.txt
    5403   32062  190989 test-data/Hamlet.txt
```

### 2. multiple text file
 
```sh
$ ./wc.exe -c -l -w  test-data/Hamlet.txt test-data/Macbeth.txt 
    3251   18164  109938 test-data/Macbeth.txt
    5403   32062  190989 test-data/Hamlet.txt
    8654   50226  300927 total
```
### 3. potected file
 
```sh
$ ./wc.exe -c -l -w  test-data/Hamlet.txt test-data/protected.txt 
wc: test-data/protected.txt open: open test-data/protected.txt: Access is denied.
    5403   32062  190989 test-data/Hamlet.txt
    5403   32062  190989 total
```    
### 4. directory
 
```sh
$ ./wc.exe -c -l -w  test-data
wc: test-data\protected.txt open: open test-data\protected.txt: Access is denied.
    5058   26894  164490 test-data\Antony and Cleopatra.txt
    3328   21689  128311 test-data\King John.txt
    3640   22765  130962 test-data\As You Like It.txt
    5132   29147  175860 test-data\Coriolanus.txt
    4188   25889  155044 test-data\Henry VIII.txt
    4045   22850  136345 test-data\Love's Labour's Lost.txt
    3066   17345  104362 test-data\The Tempest.txt
    5403   32062  190989 test-data\Hamlet.txt
    4860   27440  166231 test-data\Troiles and Cressida.txt
    3999   25943  151694 test-data\Henry IV, part 1.txt
    3515   23837  141043 test-data\Richard II.txt
    3324   21655  129756 test-data\Titus Andronicus.txt
    4846   27642  164978 test-data\King Lear.txt
    3712   22021  129668 test-data\Taming of the Shrew.txt
    4842   28775  172774 test-data\Cymbeline.txt
    3983   25874  153838 test-data\Henry VI, part 3.txt
    3587   20787  122867 test-data\Julius Caesar.txt
    3862   23569  136859 test-data\Merry Wives of Windsor.txt
    5052   31300  187096 test-data\Richard III.txt
    4157   26729  159019 test-data\Henry VI, part 2.txt
    3251   18164  109938 test-data\Macbeth.txt
    3692   22473  129009 test-data\Much Ado About Nothing.txt
    3448   22131  128207 test-data\Merchant of Venice.txt
    4032   24302  141403 test-data\All's Well That Ends Well.txt
    4164   25712  150804 test-data\Romeo and Juliet.txt
    3220   18199  106329 test-data\Two Gentlemen of Verona.txt
    3907   23071  136399 test-data\Measure for Measure.txt
    2673   16129   94013 test-data\Comedy of Errors.txt
    2809   17074  101291 test-data\Midsummer Night's Dream.txt
    4950   27784  164007 test-data\Othello.txt
    4154   27424  162053 test-data\Henry V.txt
    3577   21362  122061 test-data\Twelfth Night.txt
    4285   27747  164330 test-data\Henry IV, part 2.txt
    3444   19571  118027 test-data\Timon of Athens.txt
    3674   22811  139181 test-data\Henry VI, part 1.txt
    3301   19481  116287 test-data\Pericles.txt
  142180  857648 5085525 total
```   
### 5. user input
 
```sh
$ ./wc.exe -c -l -w
abc
def ghi jkl
^Z
       2       4      16
```   
