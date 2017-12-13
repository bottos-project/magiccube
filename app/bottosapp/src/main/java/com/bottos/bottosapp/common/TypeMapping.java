package com.bottos.bottosapp.common;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;

import java.util.HashMap;
import java.util.Map;

/**
 * bottos
 */
public class TypeMapping {

    public static String getTypeValue(String typeName) {

        Map<String, String> typeMap = new HashMap<String, String>();
        typeMap.put("DATA_AUDITING", "0");
        typeMap.put("DATA_RECRUITMENT", "1");
        typeMap.put("DATA_CLEANING", "2");

        typeMap.put("FACIAL_RECOGNITION", "0");
        typeMap.put("MACHINE_LEARNING", "1");
        typeMap.put("VOICE_INTERACTION", "2");
        typeMap.put("MACHINE_TRANSLATION", "3");

        typeMap.put("VIDEO", "0");
        typeMap.put("VOICE", "1");
        typeMap.put("TEXT", "2");

        return typeMap.get(typeName);

    }

    public static String convertDisplayString(String blockchainRetStr) {

        JSONObject jsStr = JSONObject.parseObject(blockchainRetStr);

        if (!jsStr.containsKey("items")) return blockchainRetStr; // no item found

        JSONArray jsArray = (JSONArray) jsStr.get("items");

        for (int index = 0; index < jsArray.size(); index++) {
            if (jsArray.getJSONObject(index).containsKey("applicationDomain")) {
                String appDomainValue = jsArray.getJSONObject(index).getString("applicationDomain");
                jsArray.getJSONObject(index).replace("applicationDomain", toTypeString(appDomainValue, "applicationDomain"));
            }

            if (jsArray.getJSONObject(index).containsKey("requirementType")) {
                String requirementTypeValue = jsArray.getJSONObject(index).getString("requirementType");
                jsArray.getJSONObject(index).replace("requirementType", toTypeString(requirementTypeValue, "requirementType"));
            }

            if (jsArray.getJSONObject(index).containsKey("dataType")) {
                String dataTypeValue = jsArray.getJSONObject(index).getString("dataType");
                jsArray.getJSONObject(index).replace("dataType", toTypeString(dataTypeValue, "dataType"));
            }

        }

        return jsStr.toString();
    }


    public static String toTypeString(String typeValue, String module) {

        switch (module) {
            case "applicationDomain":
                return toTypeStringAppDomain(typeValue);

            case "requirementType":
                return toTypeStringRequirementType(typeValue);

            case "dataType":
                return toTypeStringDataType(typeValue);

            default:
                return "";
        }

    }

    public static String toTypeStringAppDomain(String typeValue) {

        switch (typeValue) {
            case "0": {
                return "FACIAL_RECOGNITION";
            }
            case "1": {
                return "MACHINE_LEARNING";
            }
            case "2": {
                return "VOICE_INTERACTION";
            }
            case "3": {
                return "MACHINE_TRANSLATION";
            }

            default: {
                return "";
            }
        }
    }

    public static String toTypeStringRequirementType(String typeValue) {

        switch (typeValue) {
            case "0": {
                return "DATA_AUDITING";
            }
            case "1": {
                return "DATA_RECRUITMENT";
            }
            case "2": {
                return "DATA_CLEANING";
            }
            default: {
                return "";
            }
        }
    }


    public static String toTypeStringDataType(String typeValue) {

        switch (typeValue) {
            case "0": {
                return "VIDEO";
            }
            case "1": {
                return "VOICE";
            }
            case "2": {
                return "TEXT";
            }
            default: {
                return "";
            }
        }
    }


}
