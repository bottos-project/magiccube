pragma solidity ^0.4.2;



import "./LibInt.sol";
import "./LibString.sol";


library LibID {

    using LibInt for *;
    using LibString for *;
    using LibID for *;

    function generateID(string source) internal returns(string _id) {

        bytes32  bytesId = keccak256(source);      
        
        

        uint temp;
        uint a;
        uint b;

        _id = new string(66);

        bytes(_id)[0] = byte(0x30);
        bytes(_id)[1] = byte(0x78);

        for(uint i = 0 ;i < 32; i++) {
            
            temp = uint(bytesId[i]);

            a = temp>>4;
            b = temp&0xF;
            
            if (a > 9) {
                    bytes(_id)[2+i+i] = byte(a+0x61-0xa);
            } else {
                bytes(_id)[2+i+i] = byte(a+0x30);
            }

            if (b > 9) {
                    bytes(_id)[2+i+i+1] = byte(b+0x61-0xa);
            } else {
                bytes(_id)[2+i+i+1] = byte(b+0x30);
            }
      
        }
        

    }  
    
}