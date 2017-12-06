
pragma solidity ^0.4.2;
import "../lib/common/LibString.sol";
import "../lib/common/LibInt.sol";
library DataExchangeDeal {
    
    using LibString for *;
    using LibInt for *;
    enum ExchangeStatus{
        INVALID,
        SUBMITTED,
        DEAL_DONE,
        EXCHANGE_STATUS_MAX
    }


    struct DataExchangeRcd {
        string exchangeID;
        uint exchangeTime;
        address firstParty;
        address secondParty;
        address witnesses;
        string exchangeSignature;
        uint amout;
        string dataRequirementID;
        string dataAssetID;
        ExchangeStatus status;
    }


    function addDataExchangeRcd(DataExchangeRcd storage _self, string _exchangeID,address _firstParty, address _secondParty,  address _witnessess, string _dataRequirementID, string _dataAssetID)  internal{
        _self.exchangeID = _exchangeID;
        _self.exchangeTime = now;
        _self.firstParty = _firstParty;
        _self.secondParty = _secondParty;
        _self.witnesses = _witnessess;
        //_self.exchangeSignature = _exchangeSignature;
        //_self.amout=_amout;
        _self.dataRequirementID = _dataRequirementID;
        _self.dataAssetID = _dataAssetID;

        _self.status = ExchangeStatus.SUBMITTED;

    }

    function isDataExchangeValid(DataExchangeRcd storage _self)  internal returns (bool) {
        
        if(_self.status != ExchangeStatus.INVALID){
            return true;
        }
        else{
            return false;
        }

    }

    function isDataExchangeReadyBuy(DataExchangeRcd storage _self)  internal returns (bool) {
        
        if(_self.status == ExchangeStatus.SUBMITTED){
            return true;
        }
        else{
            return false;
        }

    }

    function setDataExchangeStatus4DealDone(DataExchangeRcd storage _self, string exchangeSignature, uint  bidMoney)  internal{
        
        _self.status = ExchangeStatus.DEAL_DONE;
        _self.exchangeSignature = exchangeSignature;
        _self.amout = bidMoney;

    }


    function toJson(DataExchangeRcd storage _self) internal returns(string _json) {

        _json = "{";

        _json = _json.concat(_self.exchangeID.toKeyValue("exchangeID"), ",");
        _json = _json.concat(_self.exchangeTime.toKeyValue("exchangeTime"), ",");

        string memory strAddr = "0x";
        strAddr = strAddr.concat(_self.firstParty.addrToAsciiString());
        _json = _json.concat(strAddr.toKeyValue("firstParty"), ",");

        strAddr = "0x";
        strAddr = strAddr.concat(_self.secondParty.addrToAsciiString());
        _json = _json.concat(strAddr.toKeyValue("secondParty"), ",");

        strAddr = "0x";
        strAddr = strAddr.concat(_self.witnesses.addrToAsciiString());
        _json = _json.concat(strAddr.toKeyValue("witnesses"), ",");

        _json = _json.concat(_self.exchangeSignature.toKeyValue("exchangeSignature"), ",");
        _json = _json.concat(_self.amout.toKeyValue("amout"), ",");
        _json = _json.concat(_self.dataRequirementID.toKeyValue("dataRequirementID"), ",");
        _json = _json.concat(_self.dataAssetID.toKeyValue("dataAssetID"), ",");
        _json = _json.concat(uint(_self.status).toKeyValue("status"));

        _json = _json.concat("}");         
		
    }


 

/*
    function aiDataAssetInfoSave(AiDataAssetInfo _self, string assetSinature, string  assetID, string dataStoreID, uint nonce, string dataPropertyDigest, string subscription, 
                              LibAIDataAsset.ApplicationDomain domain, LibAIDataAsset.AssetDataType  assetDatatype) internal {
        
        _self.aiDataModel.dataPropertyDigest = dataPropertyDigest;
        _self.aiDataModel.subscription = subscription;
        _self.aiDataModel.applicationDomain = domain;
        _self.aiDataModel.assetDataType = assetDatatype;

        LibDataAsset.dataAssetInfoSave(_self.dataAssetInfo, assetSinature, assetID, dataStoreID, nonce);

    }


    
    function getDataItem() {       
    }   
    */
}
