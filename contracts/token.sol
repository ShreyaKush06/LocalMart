// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract CampusShop {
    struct Product {
        uint256 id;
        string name;
        uint256 price;
        string description;
        string ipfsHash; // Stores product image/details JSON
        address shopOwner;
        bool isAvailable;
    }
    
    struct Order {
        uint256 orderId;
        uint256 productId;
        address student;
        uint256 quantity;
        bool isFulfilled;
        bool isPaid;
    }
    
    mapping(uint256 => Product) public products;
    mapping(uint256 => Order) public orders;
    mapping(address => bool) public registeredShops;
    
    uint256 public productCount = 0;
    uint256 public orderCount = 0;
    
    address public admin;
    
    modifier onlyAdmin() {
        require(msg.sender == admin, "Only admin can perform this action");
        _;
    }
    
    modifier onlyShopOwner() {
        require(registeredShops[msg.sender], "Only registered shops can perform this action");
        _;
    }
    
    constructor() {
        admin = msg.sender;
    }
    
    function registerShop(address shopAddress) public onlyAdmin {
        registeredShops[shopAddress] = true;
    }
    
    function addProduct(
        string memory _name,
        uint256 _price,
        string memory _description,
        string memory _ipfsHash
    ) public onlyShopOwner {
        productCount++;
        products[productCount] = Product(
            productCount,
            _name,
            _price,
            _description,
            _ipfsHash,
            msg.sender,
            true
        );
    }
    
    function placeOrder(uint256 _productId, uint256 _quantity) public {
        require(products[_productId].isAvailable, "Product not available");
        
        orderCount++;
        orders[orderCount] = Order(
            orderCount,
            _productId,
            msg.sender,
            _quantity,
            false,
            false
        );
    }
    
    function markOrderFulfilled(uint256 _orderId) public onlyShopOwner {
        uint256 productId = orders[_orderId].productId;
        require(products[productId].shopOwner == msg.sender, "Not your product");
        orders[_orderId].isFulfilled = true;
    }
    
    function confirmPayment(uint256 _orderId) public {
        require(orders[_orderId].student == msg.sender, "Not your order");
        orders[_orderId].isPaid = true;
    }
}