
pragma solidity ^0.4.2;

import "../lib/common/LibString.sol";
import "../lib/common/LibInt.sol";
import "./LibDataAsset.sol";
library LibAIDataAsset{
    
    using LibString for *;
    using LibInt for *;
    using LibDataAsset for *;
    enum ApplicationDomain{
        INDUSTRY,
        HUMAN_RECONGINIZATION,
        MEDICAL,
        //EXCHANGED,
        APPLICATION_DOMAIN_MAX
    }

    enum AssetDataType{
        TEXT,
        VIDEO,
        VOICE,
        DATATYPE_MAX
    }
 
    struct AiDataModel {
        string dataPropertyDigest;
        string description;
        ApplicationDomain applicationDomain;
        AssetDataType assetDataType;
    }


    struct AiDataAssetInfo {
        LibDataAsset.DataAssetInfo dataAssetInfo;
        AiDataModel aiDataModel;
    }  


    function aiDataAssetInfoSave(AiDataAssetInfo storage _self, address owner, string assetSinature, string  assetID, string dataStoreID, uint nonce, string dataPropertyDigest, string description, 
                              LibAIDataAsset.ApplicationDomain domain, LibAIDataAsset.AssetDataType  assetDatatype, string dataRequirementID) internal {
        
        _self.aiDataModel.dataPropertyDigest = dataPropertyDigest;
        _self.aiDataModel.description = description;
        _self.aiDataModel.applicationDomain = domain;
        _self.aiDataModel.assetDataType = assetDatatype;

        LibDataAsset.dataAssetInfoSave(_self.dataAssetInfo, owner, assetSinature, assetID, dataStoreID, nonce, dataRequirementID);

    }


    function queryAiAssetOwner(AiDataAssetInfo storage  aidataAssetInfo) internal  returns(address accountName){
        accountName = LibDataAsset.queryAssetOwner(aidataAssetInfo.dataAssetInfo);
    } 


    function addAiAuthorization(AiDataAssetInfo storage _self, address authorizationAddress) internal {
        
        LibDataAsset.addAuthorization(_self.dataAssetInfo, authorizationAddress);
        
    }

    function setAiAssetStatus(AiDataAssetInfo storage _self, LibDataAsset.DataAssetStatus status) internal {
        
        LibDataAsset.setStatus(_self.dataAssetInfo, status);
        
    }


    function toJson(AiDataAssetInfo storage _self) internal returns(string _json) {

     
        _json = _json.concat("{");

        _json = _json.concat(_self.aiDataModel.dataPropertyDigest.toKeyValue("dataPropertyDigest"), ",");
        _json = _json.concat(_self.aiDataModel.description.toKeyValue("description"), ",");
        _json = _json.concat(uint(_self.aiDataModel.applicationDomain).toKeyValue("applicationDomain"), ",");
        _json = _json.concat(uint(_self.aiDataModel.assetDataType).toKeyValue("dataPropertyDigest"), ",");

        _json = _json.concat("\"dataAssetInfo\":");

        string memory tmpJson = LibDataAsset.toJson(_self.dataAssetInfo);      
        _json = _json.concat(tmpJson);
 

        _json = _json.concat("}");     
		
    }


    function jsonParse(AiDataAssetInfo storage _self, string _strjson) internal returns (bool) {
        _self.aiDataModel.dataPropertyDigest        = _strjson.getStringValueByKey("dataPropertyDigest");
        _self.aiDataModel.description      = _strjson.getStringValueByKey("description");
        _self.aiDataModel.applicationDomain    = ApplicationDomain(_strjson.getUintValueByKey("applicationDomain"));
        _self.aiDataModel.assetDataType    = AssetDataType(_strjson.getUintValueByKey("assetDataType"));

        if ((_self.aiDataModel.applicationDomain >= ApplicationDomain.APPLICATION_DOMAIN_MAX) || (_self.aiDataModel.assetDataType >= AssetDataType.DATATYPE_MAX)){
            return false;
        }

        return LibDataAsset.jsonParse(_self.dataAssetInfo, _strjson);

    }
}
