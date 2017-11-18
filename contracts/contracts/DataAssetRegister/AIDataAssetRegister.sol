
pragma solidity ^0.4.2;

import "./LibAIDataAsset.sol";
import "../lib/common/LibID.sol";
import "../lib/common/LibString.sol";
import "../lib/common/LibInt.sol";
import "../DataExchangeManager/DataExchangeManager.sol";



contract AIDataAssetRegister {
    using LibString for *;
    using LibInt for *;
    using LibAIDataAsset for*;
    using LibDataAsset for *;
    address    owner; 

    DataExchangeManager dataExchangeAddr;
    
    mapping(string => LibAIDataAsset.AiDataAssetInfo) aiDataAssetMap;
 
    mapping(address => string[]) aiDataAssetAddrMap;



    function AIDataAssetRegister(){
        owner = msg.sender;
    }       

    function aiDataRegist(string _registInfoJson) { 

        string memory assetSignature = _registInfoJson.getStringValueByKey("assetSignature"); 
        string  memory assetID = LibID.generateID(assetSignature);
        
        
        LibAIDataAsset.AiDataAssetInfo aiDataAsset = aiDataAssetMap[assetID];     

        bool result = LibAIDataAsset.jsonParse(aiDataAsset, _registInfoJson);
        if (false == result) return;

        aiDataAsset.dataAssetInfo.assetID = assetID;

        address owner = msg.sender;
        string[] owernAllAsset = aiDataAssetAddrMap[owner];
        owernAllAsset.push(assetID);

        string memory dateRequirementID = _registInfoJson.getStringValueByKey("dataRequirementID");        
        dataExchangeAddr.addDataExchange(owner, owner, dateRequirementID, assetID);
    } 

 




    function addAiDataAssetAuthorization(string  assetID, address authorizationAddress) {
        LibAIDataAsset.AiDataAssetInfo aiDataAsset = aiDataAssetMap[assetID];

        LibAIDataAsset.addAiAuthorization(aiDataAsset, authorizationAddress);
    }

    function setAiDataAssetStatus(string  assetID, LibDataAsset.DataAssetStatus status) {
        LibAIDataAsset.AiDataAssetInfo aiDataAsset = aiDataAssetMap[assetID];

        LibAIDataAsset.setAiAssetStatus(aiDataAsset, status);
    }


    function setDataExchangeManagerAddr(DataExchangeManager addr) {
        dataExchangeAddr = addr;
    }

    function queryDataExchangeManagerAddr() constant public returns (DataExchangeManager addr) {
        addr = dataExchangeAddr;
    }
    

    function queryAiDataAssetOwner(string  assetID) constant public returns(address owner) {
       LibAIDataAsset.AiDataAssetInfo aiDataAsset = aiDataAssetMap[assetID];

        owner = LibAIDataAsset.queryAiAssetOwner(aiDataAsset);
    }

    function queryAiAssetbyID(string  assetID) constant public returns(string _json) {
        LibAIDataAsset.AiDataAssetInfo aiDataAsset = aiDataAssetMap[assetID];

        _json = LibAIDataAsset.toJson(aiDataAsset);
    }



    function queryAiAssetbyOwner(address owner) constant public returns (string _json) {
         
        string[] owernAllAsset = aiDataAssetAddrMap[owner];
        
        uint assetTotolNum = owernAllAsset.length;
        
        _json = "{";
        

        _json = _json.concat(assetTotolNum.toKeyValue("totolNum"));

        if(assetTotolNum > 0){
            _json = _json.concat(", \"items\":[");
            
            for(uint i= 0;i < assetTotolNum;i++){
                if (i>0){
                    _json = _json.concat(",");
                }
                //assetID = owernAllAsset[i];

                LibAIDataAsset.AiDataAssetInfo aiDataAsset = aiDataAssetMap[owernAllAsset[i]];
                string memory tempJson  = LibAIDataAsset.toJson(aiDataAsset);
                _json = _json.concat(tempJson);
            }

            _json = _json.concat("]}");

        }
        else {
            _json = _json.concat("}");
        }       

    }

    function queryAiAssetbyOwnerAndDataType(address owner, LibAIDataAsset.AssetDataType dataType) constant public returns (string _json) {
         
        string[] owernAllAsset = aiDataAssetAddrMap[owner];
        
        uint assetTotolNum = owernAllAsset.length;
        
        
        

        //_json = _json.concat(assetTotolNum.toKeyValue("totolNum"));
        uint coutner = 0;
        if(assetTotolNum > 0){
            string memory _jsonTmp = _jsonTmp.concat("\"items\":[");
            
            for(uint i= 0;i < assetTotolNum;i++){

                LibAIDataAsset.AiDataAssetInfo aiDataAsset = aiDataAssetMap[owernAllAsset[i]];
                if( aiDataAsset.aiDataModel.assetDataType != dataType) continue;

                if (coutner>0){
                    _jsonTmp = _jsonTmp.concat(",");
                }
                                
                string memory tempJson  = LibAIDataAsset.toJson(aiDataAsset);
                _jsonTmp = _jsonTmp.concat(tempJson);
                coutner++;
            }

            _jsonTmp = _jsonTmp.concat("]");

        }
        else {
            _json = _json.concat("}");
        }   
 
        _json = "{";
        
        if(coutner>0)    {
            _json = _json.concat(coutner.toKeyValue("totolNum"), ",");
            _json = _json.concat(_jsonTmp);
        }else {
            _json = _json.concat(coutner.toKeyValue("totolNum"));
        }

        _json = _json.concat("}");

    }

    /*
    function querybyStatus() constant {
        
    }

    function querybyDataType() constant {
        
    }
    */
    
       
          
    
   
}
