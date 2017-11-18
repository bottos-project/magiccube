
pragma solidity ^0.4.11;

import "../lib/common/LibID.sol";

contract IDTest {
    
    using LibID for *;

    address owner;

    function IDTest() {
        owner = msg.sender;
    } 

    function getStringHast(string source) constant public returns (string){
        return LibID.generateID(source);
        
    } 
}
