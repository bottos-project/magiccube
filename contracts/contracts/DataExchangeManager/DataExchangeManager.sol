
pragma solidity ^0.4.2;

import "./DataExchangeDeal.sol";
import "../DataAssetRegister/AIDataAssetRegister.sol";
import "../TokenManager/TokenManager.sol";
import "../DataRequirement/DataRequirementManager.sol";
import "../lib/common/LibID.sol";
import "../lib/common/LibString.sol";
import "../lib/common/LibInt.sol";
contract DataExchangeManager {

    //using AIDataAssetRegister for *;

    using LibString for *;
    using LibInt for *;

    AIDataAssetRegister aiDataAssetRegisterAddr;
    TokenManager tokenManagerAddr;
    DataRequirementManager dataRequirementManagerAddr;

    address    owner;
    mapping(string => DataExchangeDeal.DataExchangeRcd) dataExchangeRcdMap;

    string[] dataExchangeRcdList;

 

    function DataExchangeManager() {
        owner = msg.sender;
    }

   

    function addDataExchange(address _secondParty, address _witnessess, string _dataRequirementID, string _dataAssetID)  {
       
        string  memory  dataExchangeString = _dataRequirementID.concat(_dataAssetID);
        string  memory  dataExchangeID = LibID.generateID(dataExchangeString);

        DataExchangeDeal.DataExchangeRcd dataexchangeRcd = dataExchangeRcdMap[dataExchangeID];

        address _firstParty = dataRequirementManagerAddr.queryDataRequirementRecruiter(_dataRequirementID);     

        if (_firstParty.addrToAsciiString().equals("") ) return;   

        DataExchangeDeal.addDataExchangeRcd(dataexchangeRcd, dataExchangeID, _firstParty,  _secondParty, _witnessess,  _dataRequirementID,  _dataAssetID);

        dataExchangeRcdList.push(dataExchangeID);

        dataRequirementManagerAddr.aiAssetCount(_dataRequirementID);    
         
    }       

    function dataAssetBuy(address _firstParty, string _exchangeSignature, string _dataRequirementID, string _dataAssetID)  {
       
        string  memory dataExchangeString = _dataRequirementID.concat(_dataAssetID);
        string  memory dataExchangeID = LibID.generateID(dataExchangeString);
        DataExchangeDeal.DataExchangeRcd dataexchangeRcd = dataExchangeRcdMap[dataExchangeID];


  
        if(!DataExchangeDeal.isDataExchangeReadyBuy(dataexchangeRcd)) return; 

   

        aiDataAssetRegisterAddr.addAiDataAssetAuthorization(_dataAssetID,_firstParty);

        address  owner = aiDataAssetRegisterAddr.queryAiDataAssetOwner(_dataAssetID);

        uint  bidMoney = dataRequirementManagerAddr.queryDataRequirementBidMoney(_dataRequirementID);

        tokenManagerAddr.transfer(_firstParty, owner, bidMoney);


        DataExchangeDeal.setDataExchangeStatus4DealDone(dataexchangeRcd, _exchangeSignature, bidMoney);         
    }

    function queryOwnerbyDataAssetID(string _dataAssetID) constant public returns (address addr) {
        addr = aiDataAssetRegisterAddr.queryAiDataAssetOwner(_dataAssetID);
    }

    function queryBidMoneybyRequirementID(string _dataRequirementID) constant public returns(uint bidmoney){
        bidmoney = dataRequirementManagerAddr.queryDataRequirementBidMoney(_dataRequirementID);
    }



    function setAIDataAssetRegisterAddr(address addr) {
        aiDataAssetRegisterAddr = AIDataAssetRegister(addr);
    }

    function queryAIDataAssetRegisterAddr() constant public returns (AIDataAssetRegister addr) {
        addr = aiDataAssetRegisterAddr;
    }
    
    function setBTOTokenAddr(address addr) {
        tokenManagerAddr = TokenManager(addr);
    }

    function queryBTOTokenAddr() constant public returns (TokenManager addr) {
        addr = tokenManagerAddr;
    }


    function setDataRequirementManagerAddr(address addr) {
        dataRequirementManagerAddr = DataRequirementManager(addr);
    }

    function querytDataRequirementManagerAddr() constant public returns (DataRequirementManager addr) {
        addr = dataRequirementManagerAddr;
    }


    function queryAllDataExchange() constant public returns(string _json) {     
        
        uint totolNum = dataExchangeRcdList.length;

        _json = "{";            

        _json = _json.concat(totolNum.toKeyValue("totalNum"));

        if(totolNum > 0){
            _json = _json.concat(", \"items\":[");
            
            for(uint i= 0;i < totolNum;i++){
                if (i>0){
                    _json = _json.concat(",");
                }

                DataExchangeDeal.DataExchangeRcd dataexchangeRcd = dataExchangeRcdMap[dataExchangeRcdList[i]];

                string memory tempJson  = DataExchangeDeal.toJson(dataexchangeRcd);
                _json = _json.concat(tempJson);
            }

            _json = _json.concat("]");

        }
        else {
            //_json = _json.concat("}");
        }   

        _json = _json.concat("}");
    }


    function queryDataExchangebyDataRequirementIDAndStatus(string dataRequirementID, DataExchangeDeal.ExchangeStatus status) constant public returns(string _jsonOut) {     
        
        uint totalNum = dataExchangeRcdList.length;

        uint targetNum = 0;

        //_json = "{";            

        //_json = _json.concat(totalNum.toKeyValue("totalNum"));

        if(totalNum > 0){
            string memory  _json = ", \"items\":[";
            
            for(uint i= 0;i < totalNum;i++){

                DataExchangeDeal.DataExchangeRcd dataexchangeRcd = dataExchangeRcdMap[dataExchangeRcdList[i]];

                if (!dataexchangeRcd.dataRequirementID.equals(dataRequirementID)) continue;
                
                
                if ((status < DataExchangeDeal.ExchangeStatus.EXCHANGE_STATUS_MAX) && (dataexchangeRcd.status != status))  continue;
                

                if (targetNum>0){
                    _json = _json.concat(",");
                }                

                string memory tempJson  = DataExchangeDeal.toJson(dataexchangeRcd);
                _json = _json.concat(tempJson);

                targetNum++;
            }

            _json = _json.concat("]");

        }
  

        _jsonOut = "{";            

        _jsonOut = _jsonOut.concat(targetNum.toKeyValue("totalNum"));

        if(targetNum > 0){
            _jsonOut = _jsonOut.concat(_json);
        }

        _jsonOut = _jsonOut.concat("}");
    }


    function queryDataExchangebyAssetOwner(address owner) constant public returns(string _jsonOut) {     
        
        uint totalNum = dataExchangeRcdList.length;

        uint targetNum = 0;

        //_json = "{";            

        //_json = _json.concat(totalNum.toKeyValue("totalNum"));

        if(totalNum > 0){
            string memory  _json = ", \"items\":[";
            
            for(uint i= 0;i < totalNum;i++){

                DataExchangeDeal.DataExchangeRcd dataexchangeRcd = dataExchangeRcdMap[dataExchangeRcdList[i]];

                if (dataexchangeRcd.secondParty != owner) continue;              

                if (targetNum>0){
                    _json = _json.concat(",");
                }                

                string memory tempJson  = DataExchangeDeal.toJson(dataexchangeRcd);
                _json = _json.concat(tempJson);

                targetNum++;
            }

            _json = _json.concat("]");

        }
  

        _jsonOut = "{";            

        _jsonOut = _jsonOut.concat(targetNum.toKeyValue("totalNum"));

        if(targetNum > 0){
            _jsonOut = _jsonOut.concat(_json);
        }

        _jsonOut = _jsonOut.concat("}");
    }

    function queryDataExchangebyRequirementRecruiter(address recruiter) constant public returns(string _jsonOut) {     
        
        uint totalNum = dataExchangeRcdList.length;

        uint targetNum = 0;

        //_json = "{";            

        //_json = _json.concat(totalNum.toKeyValue("totalNum"));

        if(totalNum > 0){
            string memory  _json = ", \"items\":[";
            
            for(uint i= 0;i < totalNum;i++){

                DataExchangeDeal.DataExchangeRcd dataexchangeRcd = dataExchangeRcdMap[dataExchangeRcdList[i]];

                if (dataexchangeRcd.firstParty != recruiter) continue;              

                if (targetNum>0){
                    _json = _json.concat(",");
                }                

                string memory tempJson  = DataExchangeDeal.toJson(dataexchangeRcd);
                _json = _json.concat(tempJson);

                targetNum++;
            }

            _json = _json.concat("]");
        } 

        _jsonOut = "{";            

        _jsonOut = _jsonOut.concat(targetNum.toKeyValue("totalNum"));

        if(targetNum > 0){
            _jsonOut = _jsonOut.concat(_json);
        }

        _jsonOut = _jsonOut.concat("}");
    }

        _jsonOut = "{";            

        _jsonOut = _jsonOut.concat(targetNum.toKeyValue("totalNum"));

        if(targetNum > 0){
            _jsonOut = _jsonOut.concat(_json);
        }

        _jsonOut = _jsonOut.concat("}");
    }


     function queryDataExchangebyDataExchangeID(string dataExchangeID) constant public returns(string _json) {
         
        DataExchangeDeal.DataExchangeRcd dataexchangeRcd = dataExchangeRcdMap[dataExchangeID];

        if (dataexchangeRcd.exchangeID.equals(dataExchangeID))
        {
            _json = "{\"totalNum\":1,\"items\":[";
            string memory tempJson  = DataExchangeDeal.toJson(dataexchangeRcd);

            _json = _json.concat(tempJson);
            _json = _json.concat("]}");
        }
        else {
            _json = "{\"totalNum\":0}";
        }               
  
    }



       
          
    
   
}
