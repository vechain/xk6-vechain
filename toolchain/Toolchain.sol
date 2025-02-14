pragma solidity 0.8.19;

contract Toolchain {

    struct Todo {
        string text;
        bool completed;
    }

    event ToolchainEvent(uint256 indexed a, bytes32 indexed b, bytes32 c);

    function setBytes32(uint8 a, bytes32 b, bytes32 c) public {
        emit ToolchainEvent(a, b, c);
    }

    function payMe() public payable {
        require(msg.value > 0, "No money sent");
    }

    function balanceOf() public view returns (uint256) {
        return address(this).balance;
    }

    function randomFunc() public pure returns (uint256, address) {
        return (1, address(0));
    }

    function getTodo() public pure returns (Todo memory) {
        return Todo({
            text: "Hello, World!",
            completed: false
        });
    }
}
