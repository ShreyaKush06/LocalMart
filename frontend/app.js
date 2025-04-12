// App.js
import React, { useState, useEffect } from 'react';
import { ethers } from 'ethers';
import './App.css';

function App() {
  const [products, setProducts] = useState([]);
  const [orders, setOrders] = useState([]);
  const [account, setAccount] = useState('');
  const [contract, setContract] = useState(null);
  const [provider, setProvider] = useState(null);

  // Initialize Ethereum connection
  const initEthereum = async () => {
    if (window.ethereum) {
      try {
        const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
        setAccount(accounts[0]);
        
        const provider = new ethers.providers.Web3Provider(window.ethereum);
        setProvider(provider);
        
        // Load contract
        const contractAddress = "YOUR_CONTRACT_ADDRESS";
        const contractABI = []; // Your contract ABI
        const signer = provider.getSigner();
        const campusShop = new ethers.Contract(contractAddress, contractABI, signer);
        setContract(campusShop);
        
        // Load initial data
        loadProducts();
        loadOrders();
      } catch (error) {
        console.error(error);
      }
    } else {
      alert('Please install MetaMask!');
    }
  };

  const loadProducts = async () => {
    try {
      // Call backend API or directly to contract
      const response = await fetch('http://localhost:8080/products');
      const data = await response.json();
      setProducts(data);
    } catch (error) {
      console.error('Error loading products:', error);
    }
  };

  const loadOrders = async () => {
    // Similar implementation to load orders
  };

  const addProduct = async (product) => {
    try {
      const tx = await contract.addProduct(
        product.name,
        product.price,
        product.description,
        product.ipfsHash
      );
      await tx.wait();
      loadProducts();
    } catch (error) {
      console.error('Error adding product:', error);
    }
  };

  const placeOrder = async (productId, quantity) => {
    try {
      const tx = await contract.placeOrder(productId, quantity);
      await tx.wait();
      loadOrders();
    } catch (error) {
      console.error('Error placing order:', error);
    }
  };

  return (
    <div className="App">
      <header>
        <h1>CampusCart - Blinkit Gap Filler</h1>
        {!account ? (
          <button onClick={initEthereum}>Connect Wallet</button>
        ) : (
          <p>Connected: {account.slice(0, 6)}...{account.slice(-4)}</p>
        )}
      </header>

      <main>
        <section className="product-list">
          <h2>Products Not Available on Blinkit</h2>
          <div className="products">
            {products.map(product => (
              <div key={product.id} className="product-card">
                <h3>{product.name}</h3>
                <p>{product.description}</p>
                <p>Price: {product.price} ETH</p>
                <button onClick={() => placeOrder(product.id, 1)}>Order Now</button>
              </div>
            ))}
          </div>
        </section>

        <section className="order-section">
          <h2>Your Orders</h2>
          {/* Display orders */}
        </section>
      </main>
    </div>
  );
}

export default App;