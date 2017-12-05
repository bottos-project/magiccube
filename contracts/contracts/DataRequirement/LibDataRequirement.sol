
pragma solidity ^0.4.2;
import "../lib/common/LibString.sol";
import "../lib/common/LibInt.sol";
library LibDataRequirement {
    
    using LibString for *;
    using LibInt for *;

    enum DataReqirementStatus{
        VALID,
        INVALID,
        DATA_REQUIREMENT_STATUS_MAX
    }


    enum DataRequirementType{
        TYPE_AUDIT,
        TYPE_RECRUIT,
        TYPE_PURGE,
        DATA_REQUIREMENT_TYPE_MAX
    }

    enum DataType{
        TYPE_VIDEO,
        TYPE_VOICE,
        TYPE_TEXT,
        DATA_TYPE_MAX
    }

    enum ApplicationDomain {
        FACE_RECOGNITION,
        MACHINE_LEARNING,
        VOICE_INTERACTION, 
        MACHINE_TRANSLATION,
        APPLICATION_DOMAIN_MAX
    }
    
    struct DataRequirement{
        string dataRequirementID;
        string  dataRequirementName;
        address  recruiter;
        DataReqirementStatus status;
        uint    expirationTime;
        uint    publishedTime;
        uint    bidMoney;
        uint    nonce;
        string  requirementSignature;
        DataRequirementType requirementType;
        string  description;
        DataType dataType;
        string  specifications;
        ApplicationDomain applicationDomain;
        string dataSample1;
        string dataSample2;
        string dataSample3;
        uint   collectionNum;

    }


    function dataRequirementSave1(DataRequirement storage _self,string dataRequirementID) internal {

        _self.dataRequirementID = dataRequirementID;
        
        _self.recruiter = msg.sender;
        _self.status = DataReqirementStatus.VALID;
        

    }

    function dataRequirementSave2(DataRequirement storage _self, string  dataRequirementName, uint expirationTime, uint publishedTime, uint bidMoney,uint nonce) internal {
        _self.dataRequirementName = dataRequirementName;
        _self.expirationTime = expirationTime;
        _self.publishedTime = publishedTime;
        _self.bidMoney = bidMoney;
        _self.nonce = nonce;
    }    

    function dataRequirementSave3(DataRequirement storage _self, string  requirementSignature ,DataRequirementType requirementType,
                             string description, DataType dataType, string  specifications, ApplicationDomain applicationDomain,
                             string dataSample1, string dataSample2, string dataSample3) internal {
        
        _self.requirementSignature = requirementSignature;

        _self.requirementType = requirementType;
        _self.description = description;
        _self.dataType = dataType;
        _self.specifications = specifications;
        _self.applicationDomain = applicationDomain;
        _self.dataSample1 = dataSample1;
        _self.dataSample2 = dataSample2;
        _self.dataSample3 = dataSample3;
    }

    function jsonParse(DataRequirement storage _self, string _strjson) internal returns (bool) {

         //_self.dataRequirementID = dataRequirementID;
        
        _self.recruiter = msg.sender;
        _self.status = DataReqirementStatus.VALID;

        _self.dataRequirementName = _strjson.getStringValueByKey("dataRequirementName");
        _self.expirationTime =  _strjson.getUintValueByKey("expirationTime");
        _self.publishedTime =  _strjson.getUintValueByKey("publishedTime");
        _self.bidMoney = _strjson.getUintValueByKey("bidMoney");
        _self.nonce = _strjson.getUintValueByKey("nonce"); 
        
        _self.requirementSignature = _strjson.getStringValueByKey("requirementSignature");


        _self.requirementType = DataRequirementType(_strjson.getIntValueByKey("requirementType"));
        _self.description = _strjson.getStringValueByKey("description");
        _self.dataType = DataType(_strjson.getIntValueByKey("dataType"));
        _self.specifications = _strjson.getStringValueByKey("specifications");
        _self.applicationDomain = ApplicationDomain(_strjson.getIntValueByKey("applicationDomain"));
        _self.dataSample1 = _strjson.getStringValueByKey("dataSample1");
        _self.dataSample2 = _strjson.getStringValueByKey("dataSample2");
        _self.dataSample3 = _strjson.getStringValueByKey("dataSample3");

        if (_self.requirementType >= DataRequirementType.DATA_REQUIREMENT_TYPE_MAX) return false;
        if (_self.dataType >= DataType.DATA_TYPE_MAX) return false;
        if (_self.applicationDomain >= ApplicationDomain.APPLICATION_DOMAIN_MAX) return false;

        return true;

    }

    function aiAssetCount (DataRequirement storage _self) internal { 

        _self.collectionNum += 1;       
    } 


    function toJson(DataRequirement storage _self) internal returns(string _json) {

        _json = "{";

        _json = _json.concat(_self.dataRequirementID.toKeyValue("dataRequirementID"), ",");
        _json = _json.concat(_self.dataRequirementName.toKeyValue("dataRequirementName"), ",");
        _json = _json.concat(_self.recruiter.toKeyValue("recruiter"), ",");
        _json = _json.concat(uint(_self.status).toKeyValue("status"), ",");
        _json = _json.concat(_self.expirationTime.toKeyValue("expirationTime"), ",");
        _json = _json.concat(_self.publishedTime.toKeyValue("publishedTimename"), ",");
        _json = _json.concat(_self.bidMoney.toKeyValue("bidMoney"), ",");
        _json = _json.concat(_self.nonce.toKeyValue("nonce"), ",");
        _json = _json.concat(_self.requirementSignature.toKeyValue("requirementSignature"), ",");
        _json = _json.concat(uint(_self.requirementType).toKeyValue("requirementType"), ",");
        
        
        _json = _json.concat(_self.description.toKeyValue("description"), ",");
        _json = _json.concat(uint(_self.dataType).toKeyValue("dataType"), ",");
        _json = _json.concat(_self.specifications.toKeyValue("specifications"), ",");
         _json = _json.concat(uint(_self.applicationDomain).toKeyValue("applicationDomain"), ",");
        _json = _json.concat(_self.dataSample1.toKeyValue("dataSample1"), ",");
        _json = _json.concat(_self.dataSample2.toKeyValue("dataSample2"), ",");
        _json = _json.concat(_self.dataSample3.toKeyValue("dataSample3"), ",");
        _json = _json.concat(_self.collectionNum.toKeyValue("collectionNum"));

        _json = _json.concat("}");  
        
		
    }



    function queryDataRequirementBidMoney(DataRequirement storage  dataRequirement) returns (uint bidMoney) {
        bidMoney = dataRequirement.bidMoney;        
    }

    function queryDataRequirementRecruiter(DataRequirement storage  dataRequirement) returns (address recruiter) {
        recruiter = dataRequirement.recruiter;        
    }    
}