# EBCDIC Reader

Reads EBCDIC file and prints out content.

Default conversion table in use is found [here](http://www.flounder.com/ebcdictoascii1.htm).

Code page 277 conversion table is found [here](http://www.easymarketplace.de/cp/cp00277_Danish_Norway.pdf) 

### Usage
```bash
# Default
ebcdic-reader <ebcdic-file>

# To use the EBCDIC 277 Codepage
ebcdic-reader -277 <ebcdic-file>
```