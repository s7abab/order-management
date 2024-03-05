"use client";
import axios from "axios";
import React, { useState, useEffect } from "react";
import AddProduct from "./AddProduct";
import ProductCard from "./ProductCard";
import { useRouter } from "next/navigation";

const AddOrder = () => {
  const [modal, setModal] = useState(false);
  const [order, setOrder] = useState({
    status: "pending",
    products: [],
    total: 0,
    currencyUnit: "inr",
  });
  // navigate to home
  const router = useRouter();
  const handleNavigate = () => {
    router.push("/");
  };
  const handleModalOpen = () => {
    setModal(!modal);
  };

  // Calculate total whenever products or currencyUnit change
  useEffect(() => {
    let totalPrice = 0;
    order.products.forEach((product) => {
      totalPrice += (product.price || 0) * (product.quantity || 1);
    });

    setOrder((prevOrder) => ({
      ...prevOrder,
      total: totalPrice,
    }));
  }, [order.products, order.currencyUnit]);

  const handleAddProducts = (newProduct) => {
    setOrder((prevOrder) => ({
      ...prevOrder,
      products: [...prevOrder.products, newProduct],
    }));
    setModal(false);
  };

  const handleStatusChange = (event) => {
    const { value } = event.target;
    setOrder((prevOrder) => ({
      ...prevOrder,
      status: value,
    }));
  };

  const handleCurrencyUnitChange = (event) => {
    const { value } = event.target;
    setOrder((prevOrder) => ({
      ...prevOrder,
      currencyUnit: value,
    }));
  };

  // craete order
  const handleAddOrder = async () => {
    try {
      const res = await axios.post(`http://localhost:8080/api/v1/order`, {
        status: order.status,
        total: order.total,
        currencyUnit: order.currencyUnit,
        items: order.products,
      });
      handleNavigate();
    } catch (error) {
      console.error("Error adding order:", error);
    }
  };
  return (
    <>
      <button
        onClick={handleModalOpen}
        className="p-2 w-[130px] bg-gray-700 text-white mt-4 mx-4 rounded-md"
      >
        Add Product
      </button>
      <select
        name="status"
        id="status"
        value={order.status}
        onChange={handleStatusChange}
      >
        <option value="pending">Pending</option>
        <option value="delivered">Delivered</option>
        <option value="cancelled">Cancelled</option>
      </select>

      <select
        name="currencyUnit"
        id="currencyUnit"
        value={order.currencyUnit}
        onChange={handleCurrencyUnitChange}
      >
        <option value="usd">USD</option>
        <option value="inr">INR</option>
      </select>
      <button
        onClick={handleAddOrder}
        className="p-2 w-[130px] bg-violet-600 text-white mt-4 mx-4 rounded-md"
      >
        Submit Order
      </button>
      <div>Total: {order.total}</div>
      {modal && <AddProduct setProduct={handleAddProducts} />}
      <div className="mx-5 mt-5 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-3">
        {order.products.map((item, index) => (
          <ProductCard key={index} product={item} />
        ))}
      </div>
    </>
  );
};

export default AddOrder;
