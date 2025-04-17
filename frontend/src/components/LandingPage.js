import React, { useState } from "react";
import { useNavigate } from "react-router-dom"; // Import useNavigate
import "./LandingPage.css";

const LandingPage = () => {
  const [selectedCategory, setSelectedCategory] = useState("All");
  const navigate = useNavigate(); // Initialize navigate

  const products = [
    { id: 1, name: "Campus T-Shirt", price: "₹0", shop: "Campus Store", verified: true, category: "Clothing" },
    { id: 2, name: "Special Chai Mix", price: "₹150", shop: "Canteen", verified: false, category: "Food" },
    { id: 3, name: "Wireless Mouse", price: "₹500", shop: "Electronics Hub", verified: true, category: "Electronics" },
    { id: 4, name: "Organic Vegetables", price: "₹200", shop: "Farmers Market", verified: false, category: "Farmers" },
  ];

  const categories = ["All", "Food", "Electronics", "Farmers", "Clothing", "Stationery"];

  const filteredProducts = selectedCategory === "All"
    ? products
    : products.filter(product => product.category === selectedCategory);

  return (
    <div className="landing-page">
      {/* Top Navbar */}
      <nav className="navbar">
        <div className="navbar-left">
          <span className="logo">LocalMart</span>
        </div>
        <div className="navbar-center">
          <input
            type="text"
            className="search-bar"
            placeholder="Search products, shops..."
          />
        </div>
        <div className="navbar-right">
          <button className="auth-btn" onClick={() => navigate("/signup")}>
            Login | Sign Up
          </button>
        </div>
      </nav>

      {/* Horizontal Category Buttons */}
      <div className="category-scroll">
        {categories.map((category) => (
          <button
            key={category}
            className={`category-btn ${selectedCategory === category ? "active" : ""}`}
            onClick={() => setSelectedCategory(category)}
          >
            {category}
          </button>
        ))}
      </div>

      {/* Welcome Box */}
      <div className="welcome-box">
        <span className="emoji">👋</span>
        <h2>Welcome to LocalMart</h2>
        <p>
          Find nearby shops with real-time info. Visit only verified and trusted
          stores.
        </p>
      </div>

      {/* Blockchain Verified Sellers Box */}
      <div className="verified-box">
        <span className="emoji">✅</span>
        <h3>Blockchain Verified Sellers</h3>
        <p>
          Our verification process ensures all sellers with ✅ badge are
          blockchain authenticated.
        </p>
      </div>

      {/* Product Grid */}
      <div className="product-grid">
        {filteredProducts.map((product) => (
          <div key={product.id} className="product-card">
            <div className="product-image">
              <img
                src={`https://via.placeholder.com/150?text=${product.name}`}
                alt={product.name}
              />
            </div>
            <div className="product-info">
              <h3>{product.name}</h3>
              <p>{product.price}</p>
              <p>{product.shop}</p>
              {product.verified && <span className="verified-badge">✔ Verified</span>}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default LandingPage;