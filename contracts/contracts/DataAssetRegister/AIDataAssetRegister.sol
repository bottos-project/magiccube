
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

    DataExchangeManager public dataExchangeAddr;
    
    
    mapping(string => LibAIDataAsset.AiDataAssetInfo) aiDataAssetMap;
 
    mapping(address => string[]) aiDataAssetAddrMap;

    mapping(string => string[]) dataRequirementToAssetMap;


    function AIDataAssetRegister(){
        owner = msg.sender;
    }       

    function aiDataRegist(string _registInfoJson) returns (bool){ 

        string memory assetSignature = _registInfoJson.getStringValueByKey("assetSignature"); 
        if (assetSignature.equals("")) return false;
        string  memory assetID = LibID.generateID(assetSignature);
        
        
        LibAIDataAsset.AiDataAssetInfo aiDataAsset = aiDataAssetMap[assetID];    

        if (!aiDataAsset.dataAssetInfo.assetID.equals("")) return false;

        bool result = LibAIDataAsset.jsonParse(aiDataAsset, _registInfoJson);
        if (false == result) return false;

        aiDataAsset.dataAssetInfo.assetID = assetID;

        address owner = msg.sender;
        string[] owernAllAsset = aiDataAssetAddrMap[owner];
        owernAllAsset.push(assetID);

        string memory dateRequirementID = _registInfoJson.getStringValueByKey("dataRequirementID");        
        dataExchangeAddr.addDataExchange(owner, owner, aiDataAsset.dataAssetInfo.dataRequirementID, assetID);

        
        string[] requirementAllAsset = dataRequirementToAssetMap[aiDataAsset.dataAssetInfo.dataRequirementID];
        requirementAllAsset.push(assetID);
        return true;
        
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

    //function queryDataExchangeManagerAddr() constant public returns (DataExchangeManager addr) {
    //    addr = dataExchangeAddr;
    //}


    

    function queryAiDataAssetOwner(string  assetID) constant public returns(address owner) {
       LibAIDataAsset.AiDataAssetInfo aiDataAsset = aiDataAssetMap[assetID];

        owner = LibAIDataAsset.queryAiAssetOwner(aiDataAsset);
    }

    function queryAiAssetbyID(string  assetID) constant public returns(string _json) {
        LibAIDataAsset.AiDataAssetInfo aiDataAsset = aiDataAssetMap[assetID];

        _json = LibAIDataAsset.toJson(aiDataAsset);
    }



    
    function queryAiAssetbyOwnerAndAttribute(address owner, LibAIDataAsset.AssetDataType dataType, string dataRequirementID) constant public returns (string _json) {
         
        string[] owernAllAsset = aiDataAssetAddrMap[owner];
        
        uint assetTotalNum = owernAllAsset.length;     
        
        

        //_json = _json.concat(assetTotalNum.toKeyValue("totalNum"));
        uint coutner = 0;
        if(assetTotalNum > 0){
            string memory _jsonTmp = _jsonTmp.concat("\"items\":[");
            
            for(uint i= 0;i < assetTotalNum;i++){

                LibAIDataAsset.AiDataAssetInfo aiDataAsset = aiDataAssetMap[owernAllAsset[i]];
                if ((LibAIDataAsset.AssetDataType.DATATYPE_MAX > dataType) && (aiDataAsset.aiDataModel.assetDataType != dataType)) continue;
                if ((!dataRequirementID.equals("")) && (( !aiDataAsset.dataAssetInfo.dataRequirementID.equals(dataRequirementID)))) continue;

                if (coutner>0){
                    _jsonTmp = _jsonTmp.concat(",");
                }
                                
                string memory tempJson  = LibAIDataAsset.toJson(aiDataAsset);
                _jsonTmp = _jsonTmp.concat(tempJson);
                coutner++;
            }

            _jsonTmp = _jsonTmp.concat("]");
        }
 
        _json = "{";
        
        if(coutner>0)    {
            _json = _json.concat(coutner.toKeyValue("totalNum"), ",");
            _json = _json.concat(_jsonTmp);
        }else {
            _json = _json.concat(coutner.toKeyValue("totalNum"));
        }

        _json = _json.concat("}");

    }

    function queryAssetbyRequirementID(string dataRequirementID) constant public returns(string _json) {
         
        string[] requirementAllAsset = dataRequirementToAssetMap[dataRequirementID];
        
        uint assetTotalNum = requirementAllAsset.length;           


        uint coutner = 0;
        if(assetTotalNum > 0){
            string memory _jsonTmp = _jsonTmp.concat("\"items\":[");
            
            for(uint i= 0;i < assetTotalNum;i++){
                
                if (coutner>0){
                    _jsonTmp = _jsonTmp.concat(",");
                }

                LibAIDataAsset.AiDataAssetInfo aiDataAsset = aiDataAssetMap[requirementAllAsset[i]];
                                
                string memory tempJson  = LibAIDataAsset.toJson(aiDataAsset);
                _jsonTmp = _jsonTmp.concat(tempJson);
                coutner++;
            }

            _jsonTmp = _jsonTmp.concat("]");

        }
       
 
        _json = "{";
        
        if(coutner>0)    {
            _json = _json.concat(coutner.toKeyValue("totalNum"), ",");
            _json = _json.concat(_jsonTmp);
        }else {
            _json = _json.concat(coutner.toKeyValue("totalNum"));
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
