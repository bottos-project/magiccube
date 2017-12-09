
pragma solidity ^0.4.11;

contract TokenManager {   
    address owner;

    mapping(address => uint) balances;

    function TokenManager() {
        owner = msg.sender;
    } 

    function addTokentoAccount(address account, uint tokens) { 
        balances[account] = balances[account] + tokens;
    } 

    function transfer(address from, address to, uint tokens) {        
        if (balances[from] < tokens) return;

        balances[from] = balances[from] - tokens;
        balances[to] = balances[to] + tokens;
    }

    function getBalance(address who) constant public returns (uint tokens){
        tokens = balances[who];
    }     
}
