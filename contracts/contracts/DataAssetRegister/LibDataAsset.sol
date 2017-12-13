
pragma solidity ^0.4.2;

import "../lib/common/LibString.sol";
import "../lib/common/LibInt.sol";

library LibDataAsset {
    
    using LibString for *;
    using LibInt for *;

    enum DataAssetStatus{
        INVALID,
        VALID,
        DATA_ASSET_STATUS_MAX
    }
    
    struct DataAssetAllStatus {
        uint consumedTimes;
        DataAssetStatus lastStatus;
    }

    struct DataAssetInfo {
        address owner; 
        string  assetSignature; 
        string  assetID;      
        uint    registerTime; 
        string dataRequirementID;
        DataAssetAllStatus dataAssetAllStatus;

        string dataStoreID;
        uint   nonce;
        uint  priceLow;
        uint  priceHigh;
        string subType;
        string featureLabel1;
        string featureLabel2;
        string featureLabel3;
        uint  size;
        uint  expirationTime;

        mapping(address => uint) authorizationMap;
        address[] authorizationList;
    }  
    
    function jsonParse(DataAssetInfo storage _self, string _strjson) internal returns (bool){
        _self.owner = msg.sender;
        _self.assetSignature   = _strjson.getStringValueByKey("assetSignature");
        _self.registerTime = now;
        _self.dataRequirementID      = _strjson.getStringValueByKey("dataRequirementID");
        _self.dataAssetAllStatus.lastStatus = DataAssetStatus.VALID;
        _self.dataStoreID      = _strjson.getStringValueByKey("dataStoreID");
        _self.nonce      = _strjson.getUintValueByKey("nonce");   

        _self.priceLow      = _strjson.getUintValueByKey("priceLow");   
        _self.priceHigh      = _strjson.getUintValueByKey("priceHigh");   
        _self.subType      = _strjson.getStringValueByKey("subType");
        _self.featureLabel1      = _strjson.getStringValueByKey("featureLabel1");
        _self.featureLabel2      = _strjson.getStringValueByKey("featureLabel2");
        _self.featureLabel3      = _strjson.getStringValueByKey("featureLabel3");

        _self.size      = _strjson.getUintValueByKey("size");   
        _self.expirationTime      = _strjson.getUintValueByKey("expirationTime");  
        
        return true;
    }
    
    function queryAssetOwner(DataAssetInfo storage  dataAssetInfo) internal  returns(address accountName){
        accountName = dataAssetInfo.owner;
    }       
    
    function setAccoutName(DataAssetInfo storage dataAssetInfo, address accountName)  internal {
        dataAssetInfo.owner = accountName;
    }   
    
    function getStatus(DataAssetInfo storage dataAssetInfo) internal  returns(DataAssetStatus assetStatus){
        assetStatus = dataAssetInfo.dataAssetAllStatus.lastStatus;
    }    

    function setStatus(DataAssetInfo storage dataAssetInfo, DataAssetStatus assetStatus) internal {
        if (assetStatus >= DataAssetStatus.DATA_ASSET_STATUS_MAX){

            return;
        }

        if((dataAssetInfo.dataAssetAllStatus.lastStatus == DataAssetStatus.INVALID)
           && (assetStatus == DataAssetStatus.VALID)) {

               return;
        }

        dataAssetInfo.dataAssetAllStatus.lastStatus = assetStatus;
    }
    
    function dataAssetValid(DataAssetInfo storage dataAssetInfo) internal  returns(bool){
        return DataAssetStatus.INVALID != dataAssetInfo.dataAssetAllStatus.lastStatus;
    }
    
    function addAuthorization(DataAssetInfo storage dataAssetInfo, address authorizationAddress) internal {
        dataAssetInfo.authorizationMap[authorizationAddress] = 1;
        dataAssetInfo.authorizationList.push(authorizationAddress);

    }

    function toJson(DataAssetInfo storage _self) internal returns(string _json) {

        _json = "{";

        _json = _json.concat(_self.owner.toKeyValue("owner"), ",");
        _json = _json.concat(_self.assetSignature.toKeyValue("assetSignature"), ",");
        _json = _json.concat(_self.assetID.toKeyValue("assetID"), ",");
        _json = _json.concat(_self.registerTime.toKeyValue("registerTime"), ",");
        _json = _json.concat(_self.dataAssetAllStatus.consumedTimes.toKeyValue("consumedTimes"), ",");
        _json = _json.concat(uint(_self.dataAssetAllStatus.lastStatus).toKeyValue("lastStatus"), ",");
        _json = _json.concat(_self.dataStoreID.toKeyValue("dataStoreID"), ",");
        _json = _json.concat(_self.nonce.toKeyValue("nonce"), ",");
        _json = _json.concat(_self.dataRequirementID.toKeyValue("dataRequirementID"), ",");

        _json = _json.concat(_self.priceLow.toKeyValue("priceLow"), ",");
        _json = _json.concat(_self.priceHigh.toKeyValue("priceHigh"), ",");
        _json = _json.concat(_self.subType.toKeyValue("subType"), ",");
        _json = _json.concat(_self.featureLabel1.toKeyValue("featureLabel1"), ",");
        _json = _json.concat(_self.featureLabel2.toKeyValue("featureLabel2"), ",");
        _json = _json.concat(_self.featureLabel3.toKeyValue("featureLabel3"), ",");
        _json = _json.concat(_self.size.toKeyValue("size"), ",");
        _json = _json.concat(_self.expirationTime.toKeyValue("expirationTime"), ",");
        

        if(_self.authorizationList.length > 0){
            _json = _json.concat(_self.authorizationList.length.toKeyValue("authorizationList"),",\"elem\":[");
            for(uint loop = 0;loop < _self.authorizationList.length;loop++){
                if (loop>0)_json = _json.concat(",");
                _json = _json.concat("\"");
                _json = _json.concat(_self.authorizationList[loop].addrToAsciiString());        
                _json = _json.concat("\"");

            }
            _json = _json.concat("]");  
        }
        else{
            _json = _json.concat(uint(0).toKeyValue("authorizationList"));
        
        }

        _json = _json.concat("}");  
    }
}
