pragma solidity >=0.5.16 <=0.6.6;

import "./Token.sol";
import "./Models.sol";
import "../interfaces/ISubscription.sol";

contract Supersymmetry is ISubscription {
    event NewRequest(bytes32 requestHash, address owner, string target);
    event NewApprovedToken(address tokenAddress);
    event StatusChanged(bytes32 requestHash, Models.Status status);
    event Mint(bytes32 requestHash, address owner, uint amount);
    event Return(bytes32 requestHash, address owner, uint amount);

    bool public isWavesTokenExist;
    address nebula;
    mapping(bytes32 => Models.Request) public requests;
    mapping(bytes32 => bool) public externalRq;
    mapping(string => Models.Token) public assets;
    mapping(address => string) public assetIdByErc20;

    constructor(address _nebula) public {
        nebula = _nebula;
    }

    function data2string(bytes32 d) internal pure returns (string memory) {
        bytes memory b = new bytes(32);
        for (uint i = 0; i < 8; i++) {
            b[i] = d[i];
        }
        return string(b);
    }

    function data2address(bytes32 d) internal pure returns (address payable) {
        return address(uint160(uint(d)));
    }

    function data2rqType(bytes32 dd) internal pure returns (Models.RqType) {
        uint d = uint(dd);
        if (d == 0) return Models.RqType.Lock;
        if (d == 1) return Models.RqType.Burn;
        revert("invalid rq type");
    }

    function data2status(bytes32 dd) internal pure returns (Models.Status) {
        uint d = uint(dd);
        if (d == 0) return Models.Status.None;
        if (d == 1) return Models.Status.New;
        if (d == 2) return Models.Status.Rejected;
        if (d == 3) return Models.Status.Success;
        if (d == 4) return Models.Status.Returned;
        revert("invalid status");
    }

    function attachData(bytes32[] memory data) public {
        require(msg.sender == nebula, "access denied");
        for (uint pos = 0; pos < data.length; ) {
            bytes32 action = data[pos]; pos++;

            if (action == bytes32("mint")) {
                string memory assetId = data2string(data[pos]); pos++;
                address owner = data2address(data[pos]); pos++;
                uint amount = uint(data[pos]); pos++;
                string memory name = data2string(data[pos]); pos++;
                string memory symbol = data2string(data[pos]); pos++;
                uint8 decimals = uint8(uint(data[pos])); pos++;
                mint(assetId, owner, amount, name, symbol, decimals);
            }

            if (action == bytes32("unlock")) {
                address tokenAddress = data2address(data[pos]); pos++;
                string memory assetId = data2string(data[pos]); pos++;
                address owner = data2address(data[pos]); pos++;
                uint amount = uint(data[pos]); pos++;
                unlock(tokenAddress, assetId, owner, amount);
            }

            if (action == bytes32("unlockEth")) {
                string memory assetId = data2string(data[pos]); pos++;
                address payable owner = data2address(data[pos]); pos++;
                uint amount = uint(data[pos]); pos++;
                unlockEth(assetId, owner, amount);
            }

            if (action == bytes32("withdrawPastDueReqest")) {
                withdrawPastDueReqest(data[pos++]);
            }

            if (action == bytes32("createEthRequest")) {
                string memory target = data2string(data[pos]); pos++;
                Models.RqType rqType = data2rqType(data[pos]); pos++;
                uint amount = uint(data[pos]); pos++;
                createEthRequest(target, rqType, amount);
            }
            
            if (action == bytes32("createTokenRequest")) {
                string memory target = data2string(data[pos]); pos++;
                Models.RqType rqType = data2rqType(data[pos]); pos++;
                uint amount = uint(data[pos]); pos++;
                address tokenAddress = data2address(data[pos]); pos++;

                createTokenRequest(target, rqType, amount, tokenAddress);
            }
            
            if (action == bytes32("changeStatus")) {
                bytes32 rqHash = data[pos]; pos++;
                Models.Status status = data2status(data[pos]); pos++;

                changeStatus(rqHash, status);
            }
            
        }
    }

    function mint(string memory assetId, address owner, uint256 amount,
        string memory name, string memory symbol, uint8 decimals) internal {

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
        // externalRq[rqHash] = true;
    }

    function unlock(address tokenAddress, string memory assetId, address owner, uint256 amount) internal {
        if (assets[assetId].tokenType == Models.TokenType.None) {
            assets[assetId] = Models.Token(tokenAddress, Models.TokenType.Internal);
            assetIdByErc20[tokenAddress] = assetId;
        } else {
            require(assets[assetId].tokenType == Models.TokenType.Internal, "invalid token type");
        }

        require(Token(tokenAddress).transferFrom(address(this), owner, amount), "invalid transfer");
        // externalRq[rqHash] = true;
    }

    function unlockEth(string memory assetId, address payable owner, uint256 amount) internal {
        address tokenAddress = address(0x00);

        if (assets[assetId].tokenType == Models.TokenType.None) {
            assets[assetId] = Models.Token(tokenAddress, Models.TokenType.Internal);
            assetIdByErc20[tokenAddress] = assetId;
        } else {
            require(assets[assetId].tokenType == Models.TokenType.Internal, "invalid token type");
        }

        owner.transfer(amount);
        // externalRq[rqHash] = true;
    }

    function withdrawPastDueReqest(bytes32 requestHash) internal {
        require(requests[requestHash].status == Models.Status.New, "request not exist");
        // require(block.number > requests[requestHash].height + timeout, "rq not expire");

        if (requests[requestHash].tokenAddress == address(0x00)) {
            requests[requestHash].owner.transfer(requests[requestHash].tokenAmount);
        } else {
            require(Token(requests[requestHash].tokenAddress).transferFrom(address(this), requests[requestHash].owner, requests[requestHash].tokenAmount), "invalid balance");
        }
        requests[requestHash].status = Models.Status.Returned;
    }

    function createEthRequest(string memory target, Models.RqType rqType, uint unlockAmount) internal returns (bytes32) {
        require(rqType == Models.RqType.Lock, "rq type is not equals token type");

        uint amount = rqType == Models.RqType.Lock ? msg.value : unlockAmount;
        bytes32 requestHash = sha256(abi.encodePacked(msg.sender, target, amount, block.number));
        require(requests[requestHash].status == Models.Status.None, "request exist");

        requests[requestHash] = Models.Request(
            {status: Models.Status.New, rType: rqType, owner: msg.sender, target: target, height: block.number, tokenAmount: amount, tokenAddress: address(0x00) });
        emit NewRequest(requestHash, msg.sender, target);
        return requestHash;
    }

    function createTokenRequest(string memory target, Models.RqType rqType, uint256 amount, address tokenAddress) internal returns (bytes32) {

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

    function changeStatus(bytes32 requestHash, Models.Status status) internal {
        require(requests[requestHash].status == Models.Status.New, "status is now new");
        require(status == Models.Status.Success || status == Models.Status.Rejected, "invalid status");
        // require(block.number <= requests[requestHash].height + timeout, "rq expire");

        address tokenAddress = requests[requestHash].tokenAddress;

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