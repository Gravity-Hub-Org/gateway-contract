pragma solidity >=0.5.16 <=0.6.6;

import "./Token.sol";
import "./Models.sol";

contract Supersymmetry {
    event NewRequest(bytes32 requestHash, address owner, string target);
    event NewApprovedToken(address tokenAddress);
    event StatusChanged(bytes32 requestHash, Models.Status status);
    event Mint(bytes32 requestHash, address owner, uint amount);
    event Return(bytes32 requestHash, address owner, uint amount);

    uint8 timeout;
    address[5] public validators;
    uint8 public bftCoefficent;
    bool public isWavesTokenExist;
    address public randomContractAddress;
    mapping(bytes32 => Models.Request) public requests;
    mapping(bytes32 => bool) public externalRq;
    mapping(string => Models.Token) public assets;
    mapping(address => string) public assetIdByErc20;

    constructor(address[5] memory newValidators, uint8 newBftCoefficent, uint8 newTimeout, address newRandomContractAddress) public {
        require(newValidators[0] != address(0x00), "empty address");
        require(newValidators[1] != address(0x00), "empty address");
        require(newValidators[2] != address(0x00), "empty address");
        require(newValidators[3] != address(0x00), "empty address");
        require(newValidators[4] != address(0x00), "empty address");

        timeout = newTimeout;
        bftCoefficent = newBftCoefficent;
        validators = newValidators;
        randomContractAddress = newRandomContractAddress;
    }

    function _requireValidSign(bytes32 message, uint8[5] memory v,
        bytes32[5] memory r, bytes32[5] memory s) internal {

        int count = 0;
        for(uint i = 0; i < 5; i++) {
            count += ecrecover(keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", message)),
                v[i], r[i], s[i]) == validators[i] ? 1 : 0;
        }

        require(count > bftCoefficent, "invalid signs");
    }

    function mint(bytes32 rqHash, string memory assetId, address owner, uint256 amount,
        uint8[5] memory v, bytes32[5] memory r, bytes32[5] memory s, string memory name, string memory symbol, uint8 decimals) public {

        bytes32 message = keccak256(abi.encodePacked(rqHash, assetId, owner, amount, name, symbol, decimals, Models.TokenType.External));
        _requireValidSign(message, v, r, s);

        address tokenAddress;
        if (assets[assetId].tokenType == Models.TokenType.None) {
            tokenAddress = address(new Token(name, symbol, decimals));
            assets[assetId] = Models.Token(tokenAddress, Models.TokenType.External);
            assetIdByErc20[tokenAddress] = assetId;
        } else {
            require(assets[assetId].tokenType == Models.TokenType.External, "invalid token type");
            tokenAddress = assets[assetId].tokenAddress;
        }

        require(Token(tokenAddress).mint(owner, amount), "invalid mint");
        externalRq[rqHash] = true;
    }

    function unlock(bytes32 rqHash, address tokenAddress, string memory assetId, address owner, uint256 amount,
        uint8[5] memory v, bytes32[5] memory r, bytes32[5] memory s) public {

        int count = 0;
        for(uint i = 0; i < 5; i++) {
            count += ecrecover(keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n33", abi.encodePacked(rqHash,
                rqHash, assetId, owner, amount, tokenAddress, Models.TokenType.Internal))),
                v[i], r[i], s[i]) == validators[i] ? 1 : 0;
        }

        if (assets[assetId].tokenType == Models.TokenType.None) {
            assets[assetId] = Models.Token(tokenAddress, Models.TokenType.Internal);
            assetIdByErc20[tokenAddress] = assetId;
        } else {
            require(assets[assetId].tokenType == Models.TokenType.Internal, "invalid token type");
        }

        require(Token(tokenAddress).transferFrom(address(this), owner, amount), "invalid transfer");
        externalRq[rqHash] = true;
    }

    function unlockEth(bytes32 rqHash, string memory assetId, address payable owner, uint256 amount,
        uint8[5] memory v, bytes32[5] memory r, bytes32[5] memory s) public payable {

        address tokenAddress = address(0x00);
        int count = 0;
        for(uint i = 0; i < 5; i++) {
            count += ecrecover(keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n33", abi.encodePacked(rqHash,
                rqHash, assetId, owner, amount, tokenAddress, Models.TokenType.Internal))),
                v[i], r[i], s[i]) == validators[i] ? 1 : 0;
        }

        if (assets[assetId].tokenType == Models.TokenType.None) {
            assets[assetId] = Models.Token(tokenAddress, Models.TokenType.Internal);
            assetIdByErc20[tokenAddress] = assetId;
        } else {
            require(assets[assetId].tokenType == Models.TokenType.Internal, "invalid token type");
        }

        owner.transfer(amount);
        externalRq[rqHash] = true;
    }


    function withdrawPastDueReqest(bytes32 requestHash) public {
        require(requests[requestHash].status == Models.Status.New, "request not exist");
        require(block.number > requests[requestHash].height + timeout, "rq not expire");

        if (requests[requestHash].tokenAddress == address(0x00)) {
            requests[requestHash].owner.transfer(requests[requestHash].tokenAmount);
        } else {
            require(Token(requests[requestHash].tokenAddress).transferFrom(address(this), requests[requestHash].owner, requests[requestHash].tokenAmount), "invalid balance");
        }
        requests[requestHash].status = Models.Status.Returned;
    }

    function createEthRequest(string memory target, Models.RqType rqType, uint unlockAmount) public payable returns (bytes32) {
        require(rqType == Models.RqType.Lock, "rq type is not equals token type");

        uint amount = rqType == Models.RqType.Lock ? msg.value : unlockAmount;
        bytes32 requestHash = sha256(abi.encodePacked(msg.sender, target, amount, block.number));
        require(requests[requestHash].status == Models.Status.None, "request exist");

        requests[requestHash] = Models.Request(
            {status: Models.Status.New, rType: rqType, owner: msg.sender, target: target, height: block.number, tokenAmount: amount, tokenAddress: address(0x00) });
        emit NewRequest(requestHash, msg.sender, target);
        return requestHash;
    }

    function createTokenRequest(string memory target, Models.RqType rqType, uint256 amount, address tokenAddress) public returns (bytes32) {

        require(tokenAddress != address(0x00), "invalid tokenAddress");
        require(amount > 0, "value less or equals 0");

        if (rqType == Models.RqType.Burn || rqType == Models.RqType.Lock) {
            require(Token(tokenAddress).transferFrom(msg.sender, address(this), amount), "invalid balance");
        }

        bytes32 requestHash = sha256(abi.encodePacked(msg.sender, target, amount, block.number));
        require(requests[requestHash].status == Models.Status.None, "request exist");

        requests[requestHash] = Models.Request(
            {status: Models.Status.New, rType: rqType, owner: msg.sender, target: target, height: block.number, tokenAmount: amount, tokenAddress: tokenAddress });
        emit NewRequest(requestHash, msg.sender, target);
        return requestHash;
    }

    function changeStatus(bytes32 requestHash, uint8[5] memory v, bytes32[5] memory r, bytes32[5] memory s, Models.Status status) public {
        require(requests[requestHash].status == Models.Status.New, "status is now new");
        require(status == Models.Status.Success || status == Models.Status.Rejected, "invalid status");
        require(block.number <= requests[requestHash].height + timeout, "rq expire");

        address tokenAddress = requests[requestHash].tokenAddress;
        int count = 0;
        for(uint i = 0; i < 5; i++) {
            count += ecrecover(keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n33", abi.encodePacked(requestHash, uint8(status)))),
                v[i], r[i], s[i]) == validators[i] ? 1 : 0;
        }

        require(count >= bftCoefficent, "admins vote count is less bftCoefficent");

        if (status == Models.Status.Rejected) {
             if (requests[requestHash].rType == Models.RqType.Burn) {
                 require(Token(tokenAddress).transferFrom(address(this), requests[requestHash].owner, requests[requestHash].tokenAmount), "invalid balance");
        } else if (requests[requestHash].rType == Models.RqType.Lock) {
                if (tokenAddress == address(0x00)) {
                    requests[requestHash].owner.transfer(requests[requestHash].tokenAmount);
                } else {
                    require(Token(tokenAddress).transferFrom(address(this), requests[requestHash].owner, requests[requestHash].tokenAmount), "invalid balance");
                }
             }
         }
        requests[requestHash].status = status;
        emit StatusChanged(requestHash, status);
    }
}