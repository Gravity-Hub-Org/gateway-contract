pragma solidity >=0.5.16 <=0.6.6;

library Models {
    enum Status {
        None,
        New,
        Rejected,
        Success,
        Returned
    }

    enum RqType {
        Lock,
        Unlock,
        Mint,
        Burn
    }
    enum TokenType {
        NativeToken,
        InputToken
    }

    struct Token {
        string assetId;
        TokenType tokenType;
        Status status;
        uint8 decimals;
    }

    struct Request {
        Status status;
        RqType rType;
        address payable owner;
        uint256 height;
        string target;
        uint256 tokenAmount;
        address tokenAddress;
        string targetRqId;
    }
}