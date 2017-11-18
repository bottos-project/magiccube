pragma solidity ^0.4.2;

import "./LibString.sol";
import "./LibInt.sol";

library LibJSON {
    
    using LibJSON for *;  
    using LibString for *;  
    using LibInt for *;

    function getObjectValueByKey(string _self, string _key) internal returns (string _ret) {
        int pos = -1;
        uint searchStart = 0;
        while (true) {
            pos = _self.indexOf("\"".concat(_key, "\""), searchStart);
            if (pos == -1) {
                pos = _self.indexOf("'".concat(_key, "'"), searchStart);
                if (pos == -1) {
                    return;
                }
            }

            pos += int(bytes(_key).length+2);
            // pos 起始为：{
            bool colon = false;
            while (uint(pos) < bytes(_self).length) {
                if (bytes(_self)[uint(pos)] == ' ' || bytes(_self)[uint(pos)] == '\t' 
                    || bytes(_self)[uint(pos)] == '\r' || bytes(_self)[uint(pos)] == '\n') {
                    pos++;
                } else if (bytes(_self)[uint(pos)] == ':') {
                    pos++;
                    colon = true;
                    break;
                } else {
                    break;
                }
            }

            if(uint(pos) == bytes(_self).length) {
                return;
            }

            if (colon) {
                break;
            } else {
                searchStart = uint(pos);
            }
        }

        int start = _self.indexOf("{", uint(pos));
        if (start == -1) {
            return;
        }
        //start += 1;
        
        int end = _self.indexOf("}", uint(pos));
        if (end == -1) {
            return;
        }
        end +=1 ;
        _ret = _self.substr(uint(start), uint(end-start));
    }

    function getIntArrayValueByKey(string _self, string _key, uint[] storage _array) internal {
         for (uint i=0; i<10; ++i) {
            //delete _array[i];
            _array[i] = i;
        }
        //_array.length = 0;

        /*
        int pos = -1;
        uint searchStart = 0;
        while (true) {
            pos = _self.indexOf("\"".concat(_key, "\""), searchStart);
            if (pos == -1) {
                pos = _self.indexOf("'".concat(_key, "'"), searchStart);
                if (pos == -1) {
                    return;
                }
            }

            pos += int(bytes(_key).length+2);

            bool colon = false;
            while (uint(pos) < bytes(_self).length) {
                if (bytes(_self)[uint(pos)] == ' ' || bytes(_self)[uint(pos)] == '\t' 
                    || bytes(_self)[uint(pos)] == '\r' || bytes(_self)[uint(pos)] == '\n') {
                    pos++;
                } else if (bytes(_self)[uint(pos)] == ':') {
                    pos++;
                    colon = true;
                    break;
                } else {
                    break;
                }
            }

            if(uint(pos) == bytes(_self).length) {
                return;
            }

            if (colon) {
                break;
            } else {
                searchStart = uint(pos);
            }
        }

        int start = _self.indexOf("[", uint(pos));
        if (start == -1) {
            return;
        }
        start += 1;
        
        int end = _self.indexOf("]", uint(pos));
        if (end == -1) {
            return;
        }

        string memory vals = _self.substr(uint(start), uint(end-start)).trim(" \t\r\n");

        if (bytes(vals).length == 0) {
            return;
        } */
        
        

        // string[] memory _strArray ;
        // vals.split(",", _strArray);

        // for (uint i=0; i<_strArray.length; ++i) {
        //     _array[i] = _strArray[i].trim(" \t\r\n");
        //     _array[i] = _strArray[i].trim("'\"");
        // }
    }
}