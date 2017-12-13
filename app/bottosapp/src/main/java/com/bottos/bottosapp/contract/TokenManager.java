package com.bottos.bottosapp.contract;

import java.math.BigInteger;
import java.util.Arrays;
import java.util.Collections;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Function;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.generated.Uint256;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.RemoteCall;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.tx.Contract;
import org.web3j.tx.TransactionManager;

/**
 * <p>Auto generated code.
 * <p><strong>Do not modify!</strong>
 * <p>Please use the <a href="https://docs.web3j.io/command_line.html">web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the 
 * <a href="https://github.com/web3j/web3j/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 3.0.1.
 */
public final class TokenManager extends Contract {
    private static final String BINARY = "6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a031990911617905561019a8061003b6000396000f3006060604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663beabacc8811461005b578063f8b2cb4f14610085578063ff895030146100b6575b600080fd5b341561006657600080fd5b610083600160a060020a03600435811690602435166044356100d8565b005b341561009057600080fd5b6100a4600160a060020a0360043516610131565b60405190815260200160405180910390f35b34156100c157600080fd5b610083600160a060020a036004351660243561014c565b600160a060020a038316600090815260016020526040902054819010156100fe5761012c565b600160a060020a03808416600090815260016020526040808220805485900390559184168152208054820190555b505050565b600160a060020a031660009081526001602052604090205490565b600160a060020a039091166000908152600160205260409020805490910190555600a165627a7a72305820d071b9a118d85ae01f0769f11bffc8b2d05eb8dfab5c54c95198a9b5ee6720b30029";

    public TokenManager(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    private TokenManager(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public RemoteCall<TransactionReceipt> transfer(String from, String to, BigInteger tokens) {
        Function function = new Function(
                "transfer", 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(from), 
                new org.web3j.abi.datatypes.Address(to), 
                new org.web3j.abi.datatypes.generated.Uint256(tokens)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<BigInteger> getBalance(String who) {
        Function function = new Function("getBalance", 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(who)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteCall<TransactionReceipt> addTokentoAccount(String account, BigInteger tokens) {
        Function function = new Function(
                "addTokentoAccount", 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(account), 
                new org.web3j.abi.datatypes.generated.Uint256(tokens)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public static RemoteCall<TokenManager> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(TokenManager.class, web3j, credentials, gasPrice, gasLimit, BINARY, "");
    }

    public static RemoteCall<TokenManager> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(TokenManager.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, "");
    }

    public static TokenManager load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new TokenManager(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    public static TokenManager load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new TokenManager(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }
}
