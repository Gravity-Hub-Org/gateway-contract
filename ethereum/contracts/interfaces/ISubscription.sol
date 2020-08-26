pragma solidity >=0.4.21 <0.7.0;

interface ISubscription {
    function attachData(bytes32[] calldata data) external;
}