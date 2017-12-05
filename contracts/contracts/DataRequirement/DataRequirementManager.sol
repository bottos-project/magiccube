
pragma solidity ^0.4.2;
import "../lib/common/LibString.sol";
import "../lib/common/LibInt.sol";
import "./LibDataRequirement.sol";
import "../lib/common/LibID.sol";

contract DataRequirementManager {
    using LibString for *;
    using LibInt for *;
    
    using LibDataRequirement for *;

    address    owner;
    mapping(string => LibDataRequirement.DataRequirement) dataRequirementMap;

    string [] dataRequirementList;


    function DataRequirementManager(){
        owner = msg.sender;
    }    


    function createDataRequirement (string  _requirementJson) returns (bool) { 
              
        string memory requirementSignature = _requirementJson.getStringValueByKey("requirementSignature");

        if (requirementSignature.equals("")) return false;

        string  memory dataRequirementID = LibID.generateID(requirementSignature);
        

        LibDataRequirement.DataRequirement dataRequirement = dataRequirementMap[dataRequirementID];

        if (!dataRequirement.dataRequirementID.equals("")) return false;

        bool result = LibDataRequirement.jsonParse(dataRequirement, _requirementJson);
        if (false == result) {
            delete dataRequirementMap[dataRequirementID];
            return false;
        }

        dataRequirement.dataRequirementID = dataRequirementID;        

        dataRequirementList.push(dataRequirementID);

        return true;
    } 

    function aiAssetCount (string  dataRequirementID) { 

        LibDataRequirement.DataRequirement dataRequirement = dataRequirementMap[dataRequirementID];

        if (!dataRequirement.dataRequirementID.equals(dataRequirementID)) {
            return;
        }

        LibDataRequirement.aiAssetCount(dataRequirement);        
    } 

    function queryDataRequirementbyID(string _id) constant public returns(string _json) {        
        
        LibDataRequirement.DataRequirement dataRequirement = dataRequirementMap[_id];
        if (!dataRequirement.dataRequirementID.equals(""))  {
             _json = _json.concat("{");
             _json = _json.concat(uint(1).toKeyValue("totalNum"), ",");
             _json = _json.concat("\"items\":[");

            string memory tmpJson = LibDataRequirement.toJson(dataRequirement);      
            _json = _json.concat(tmpJson);

            _json = _json.concat("]}");
        }
        else
        {
            _json = _json.concat("{");
            _json = _json.concat(uint(0).toKeyValue("totalNum"));
            _json = _json.concat("}");
        }
    }



    function queryDataRequirementBidMoney(string dataRequirementID) constant public returns (uint bidMoney) {
        LibDataRequirement.DataRequirement dataRequirement = dataRequirementMap[dataRequirementID];

        bidMoney = LibDataRequirement.queryDataRequirementBidMoney(dataRequirement);
    }

    function queryDataRequirementRecruiter(string dataRequirementID) constant public returns (address recruiter) {
        LibDataRequirement.DataRequirement dataRequirement = dataRequirementMap[dataRequirementID];

        recruiter = LibDataRequirement.queryDataRequirementRecruiter(dataRequirement);
    }
   

    function queryDataRequirementbyOwner(address recruiter) constant public returns(string _json) {     
        
              
        string memory _jsonTmp;

        uint counter = 0;
        for (uint i=0; i<dataRequirementList.length; ++i) {
            //LibDataRequirement.DataRequirement tmpData = dataRequirementList[i];
            LibDataRequirement.DataRequirement dataRequirement = dataRequirementMap[dataRequirementList[i]];
            if (dataRequirement.recruiter != recruiter) {
                continue;
            }

            if (counter > 0){
                _jsonTmp = _jsonTmp.concat(",");
            }

            
            string memory tmpLoopJson = LibDataRequirement.toJson(dataRequirement); 

            _jsonTmp = _jsonTmp.concat(tmpLoopJson);

            counter++;
        }
   

        _json = "{"; 

        _json = _json.concat(counter.toKeyValue("totalNum"));

        if(counter>0){
            _json = _json.concat(", \"items\":[");

            _json = _json.concat(_jsonTmp);

            _json = _json.concat("]");
        }
        else{
            
        }

        _json = _json.concat("}");
    }

    function queryDataRequirementbyType(LibDataRequirement.DataRequirementType requirementType) constant public returns(string _json) {     
        
              
        string memory _jsonTmp;

        uint counter = 0;
        for (uint i=0; i<dataRequirementList.length; ++i) {
            //LibDataRequirement.DataRequirement tmpData = dataRequirementList[i];
            LibDataRequirement.DataRequirement dataRequirement = dataRequirementMap[dataRequirementList[i]];
            if (dataRequirement.requirementType != requirementType) {
                continue;
            }

            if (counter > 0){
                _jsonTmp = _jsonTmp.concat(",");
            }

            
            string memory tmpLoopJson = LibDataRequirement.toJson(dataRequirement); 

            _jsonTmp = _jsonTmp.concat(tmpLoopJson);

            counter++;
        }
   

        _json = "{"; 

        _json = _json.concat(counter.toKeyValue("totalNum"));

        if(counter>0){
            _json = _json.concat(", \"items\":[");

            _json = _json.concat(_jsonTmp);

            _json = _json.concat("]");
        }
        else{
            
        }

        _json = _json.concat("}");
    }

    function queryDataRequirementbyDataType(LibDataRequirement.DataType dataType) constant public returns(string _json) {     
        
              
        string memory _jsonTmp;

        uint counter = 0;
        for (uint i=0; i<dataRequirementList.length; ++i) {
            //LibDataRequirement.DataRequirement tmpData = dataRequirementList[i];
            LibDataRequirement.DataRequirement dataRequirement = dataRequirementMap[dataRequirementList[i]];
            if (dataRequirement.dataType != dataType) {
                continue;
            }

            if (counter > 0){
                _jsonTmp = _jsonTmp.concat(",");
            }

            
            string memory tmpLoopJson = LibDataRequirement.toJson(dataRequirement); 

            _jsonTmp = _jsonTmp.concat(tmpLoopJson);

            counter++;
        }
   

        _json = "{"; 

        _json = _json.concat(counter.toKeyValue("totalNum"));

        if(counter>0){
            _json = _json.concat(", \"items\":[");

            _json = _json.concat(_jsonTmp);

            _json = _json.concat("]");
        }
        else{
            
        }

        _json = _json.concat("}");
    }

    function queryDataRequirementbyApplicationDomain(LibDataRequirement.ApplicationDomain applicationDomain) constant public returns(string _json) {     
        
              
        string memory _jsonTmp;

        uint counter = 0;
        for (uint i=0; i<dataRequirementList.length; ++i) {
            //LibDataRequirement.DataRequirement tmpData = dataRequirementList[i];
            LibDataRequirement.DataRequirement dataRequirement = dataRequirementMap[dataRequirementList[i]];
            if (dataRequirement.applicationDomain != applicationDomain) {
                continue;
            }

            if (counter > 0){
                _jsonTmp = _jsonTmp.concat(",");
            }

            
            string memory tmpLoopJson = LibDataRequirement.toJson(dataRequirement); 

            _jsonTmp = _jsonTmp.concat(tmpLoopJson);

            counter++;
        }
   

        _json = "{"; 

        _json = _json.concat(counter.toKeyValue("totalNum"));

        if(counter>0){
            _json = _json.concat(", \"items\":[");

            _json = _json.concat(_jsonTmp);

            _json = _json.concat("]");
        }
        else{
            
        }

        _json = _json.concat("}");
    }

    function queryAllDataRequirement() constant public returns(string _json) {     
        
        uint totalNum = dataRequirementList.length;

        _json = "{";            

        _json = _json.concat(totalNum.toKeyValue("totalNum"));

        if(totalNum > 0){
            _json = _json.concat(", \"items\":[");
            
            for(uint i= 0;i < totalNum;i++){
                if (i>0){
                    _json = _json.concat(",");
                }

                LibDataRequirement.DataRequirement dataRequirement = dataRequirementMap[dataRequirementList[i]];

                string memory tempJson  = LibDataRequirement.toJson(dataRequirement);
                _json = _json.concat(tempJson);
            }

            _json = _json.concat("]");

        }
        else {
            //_json = _json.concat("}");
        }   

        _json = _json.concat("}");
    }

}


