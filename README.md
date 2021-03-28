# Gemini Fortune Bot

```
                                    :=+*#######**+-.                                      
                               .:-=#%%%%%%%%%%%%%%#%#+=-:.                                
                           .=+*********++*+*=*++*#*####*##*=.                             
                         .+************+==*+--==#*#******####+.                           
                         +#***####*##*+%%==+==#%=**##*##*##%%#*-                          
                        .#*#####*##*##+*%=+==-=++####%##########=                         
                        :*#**#####%%%%%=-=+%:=-=#%%%%%###%%#%##%%-                        
                       .***#####%%%%%%%%%%#*#%%%%%%%%%%%#*######%*                        
                       *##*###%%%%%#**#%#+*#*++#%##*#%%%%%**#**##+                        
                      :####%#%%%#=::::=+========+=::::=#%%%#%**##-                        
                       *%#%%#%%*:::--==============--:::*%%#%%*#*                         
                       *#%%#%%%:-=###%%#*++===+*#%%##*=-:%%%*%%*#                         
                       *#%%%%%=:==++++*#%%#--#%%%*++++==-=%%%%#*=                         
                       *%%%%+=::-+:..-.:#*-:::+#:.-..:+=::+*%%%%.                         
                       +%%==:::::::-+#%%+::--::=#%%+=:::::::+=%=                          
                      .  ::=:-:::::======::-:::-=====:::::--+--  .                        
                   :-*####+-:::::::======::-:::-=====-::::---+*##**--                     
                  .=**%%%%@:-+.:::======--:+=:--======::::*-.%%%%%**=.                    
                 =#**#%%#%#==*+::-======+##%###+======-::=*+=#%#%%#*#*=                   
                -*#**##***==++-::=====-==+=%%=+==-=====::-*=+****###***-                  
                -********-.*%*-=:====#%%#%%%#%%#%%#+===:==*%@**********=                  
               :**#***#**.=++%#=*+=-+#%%%%%%%%%%%%#+-==*=#%%%#***#***#**:                 
              .**#*******:%#:%%+=##--%#--=+**+=--*%=-*#++%%%%%********#**:                
             :*#*#*******+@+.%%%#*%#+##=#%%%%%@#=***%#*#%%%%%%#*******###*.               
             *#*#####****+%##%%%-%%%#%@::=+%%*=::#%#%%%=@%%%%%*******###*#*.              
             +*###%###%#***%%%%%:#%%%%+:++%%%%++:+%%%%#-%%%%%***#%###%###**               
             :-:++*#%%%%##***#%#=+#%%%--+%%%%%%+==%%%*::#%#***#%%%%%#*=+:-:               
                  --+#%%%%%##**#+-:#%%%%%%%%%%%%%%%%+:::***##%%%%%#+--                    
                  :=+%%#%%%%%%#*=:::+%%%%%%%%%%%%%%-::::+#%%%%%%#%%#=.                    
                .*%%%##%#%%%%#%#*=:::-*%%%%%%%%%%#-:::::*%%%%#%##%%%%%*=                  
                +*#%%%%%***#%%%##*-:::-*@%%%%%%%%-::::++*%#***%%%%%%%%%%%.                
           .:-+#%#*###%%%%%#*%%%%#*-:::=+%%%%%%#=::::+%%#*%%#%%%%%%%%%%%%%*=:.            
      .=++##****##**%#*%%%%####%%%#*::::=+@%%%%=-::+#%%#*#%%%%%%%%#%%%%%%%%%##**+==:      
    -*#%%########%%****##%%%%%**##%%*:.:==*%%%==::#%%**#%%%%%%####%%%%%%%%%%######*#*=    
   =##%##*###**##%%%#*****%%%%%%##%%#+::-==+*===+*##%%#%%%%%*#**##%%%%%%%%#%%%####%#*#+   
  .#%%#***####-+#+#%%%#***##*+%++#-+*%=:.=====.#+*##%=*=#%%#*##*%%%%%%%%%%-%%%%###*%%#*+  
  .#*%##*#%####:++*+%+%====*#.#+.%-+-#+-.:-=== #=*#%--#=*+=+-==*++#.#++#** #####***##%%#  
  =###*#*######=:=*++:@:-=-##+.=:-++ +==-+.=+* =:-#%=:%--:%=:+ =+-*:*:++ % ##***#****%%@. 
  -*##*##*#%%#%++#+*%+*++#%##*=***%%=***+#+#%#=*++%%@*=-*=**+*+*#*#+%++#++=%%%%#####*#%%. 
  .*#**#%#%#%%#%%%%%%%%%%%%%#*#***#%%%*%#%%%%%*#@%%%#*##*#%%%%%%%%%%%%##%##%%%%*#**#**%+  
   :*##%%%####%%%%##%#%%%%%%%%#***##%###%%#%%##%%%%#***##%%%%%%%%%%%##%%#*##*#%#######%:  
      :=#%%##%%%%%%%*###%%%%%%%%#***##%%%%%%%%%%%#***##%%%%%%%%%%###%%%%%%##*#%%###%*=:   
         .:-=+#%%##*%###%%#%%%%%%%##**#%%%%%%%%#**#%#%%%%%%%%%####%%#%%%%%%%##+-:::.      
               .....:--+***#%%%#%%%%%**###%%##****#%%%%%%#%%#######*+===+=-.              
                              :-==+*******+++***#**=--------:..                           

```


## Intro

This bot acts as a `fortune teller machine` similar to Zoltar. If configured correctly, the bot can provide a user with a fortune after a transaction of any size is sent to the bots cryptocurrency address.

The fortune manifests itself as a gemlog on a [Gemini website](https://gemini.circumlunar.space/) and within the associated ATOM feed.

## Installation

    git clone https://github.com/t-900-a/gemini-fortune-bot
    go build .

A program called `gemini-fortune-bot` will end up in your `current` directory.

## Usage

    gemini-fortune-bot [/path/to/fortune/cookie/file] [website uri] [tx hash] [payment uri] [payment view key]
    fortune -h|--help

If you don't specify a fortune cookie file path (see below), _fortune_
defaults to the contents of the `FORTUNE_FILE` environment variable. If
neither the argument nor the `FORTUNE_FILE` variable is present, _fortune_
aborts.

For the bot to run only after a small payment has been made see below:

    monero-wallet-rpc --wallet-file ~/mywallet --prompt-for-password \
    --tx-notify "/usr/bin/gemini-fortune-bot /var/fortunes/fortunes monero:donate.getmonero.org %s"


## Fortune Cookie File Format

A fortune cookie file is a text file full of quotes. The format is simple:
The file consists of paragraphs separated by lines containing a single '%'
character. For example::

    A little caution outflanks a large cavalry.
        -- Bismarck
    %
    A little retrospection shows that although many fine, useful software
    systems have been designed by committees and built as part of multipart
    projects, those software systems that have excited passionate fans are
    those that are the products of one or a few designing minds, great
    designers. Consider Unix, APL, Pascal, Modula, the Smalltalk interface,
    even Fortran; and contrast them with Cobol, PL/I, Algol, MVS/370, and
    MS-DOS.
        -- Fred Brooks, Jr.
    %
    A man is not old until regrets take the place of dreams.
        -- John Barrymore

You're more than welcome to _this_ fortune cookie file. It's over here:
<https://github.com/bmc/fortunes>.

## License and Copyright

See the accompanying License file.
