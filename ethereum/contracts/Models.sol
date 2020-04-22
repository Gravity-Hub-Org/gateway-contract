pragma solidity >=0.5.16 <=0.6.6;

library Models {
    enum Status {
        None,
        New,
        Rejected,
        Success,
        Returned
    }

    enum TokenType {
        None,
        Internal,
        External
    }

    enum RqType {
        Lock,
        Burn
    }

    struct Token {
        address tokenAddress;
        TokenType tokenType;
    }

    struct Request {
        Status status;
        RqType rType;
        address payable owner;
        uint256 height;
        string target;
        uint256 tokenAmount;
        address tokenAddress;
    }
}