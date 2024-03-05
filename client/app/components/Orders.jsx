"use client";
import axios from "axios";
import { useRouter } from "next/navigation";
import { useState, useEffect } from "react";

const Orders = () => {
  const [orders, setOrders] = useState([]);
  const router = useRouter();

  const handleRoute = () => {
    router.push("/addorder");
  };

  const fetchOrders = async () => {
    const res = await axios.get(
      `http://localhost:8080/api/v1/orders?page=1&pageSize=10&sortBy=total&sortOrder=DESC&status=PENDING`
    );
    setOrders(res.data);
  };
  useEffect(() => {
    fetchOrders();
  }, []);
  return (
    <div className="w-screen h-screen p-4 bg-white rounded-lg shadow-md">
      <button
        onClick={handleRoute}
        className="p-2 w-[100px] bg-gray-700 text-white m-4 mx-4 rounded-md"
      >
        Add Order
      </button>
      <table className="w-full table-auto">
        <thead>
          <tr className="bg-gray-200">
            <th className="px-4 py-2">ID</th>
            <th className="px-4 py-2">Status</th>
            <th className="px-4 py-2">Total</th>
            <th className="px-4 py-2">Currency Unit</th>
            <th className="px-4 py-2">Actions</th>
          </tr>
        </thead>
        <tbody>
          {orders?.map((order, index) => (
            <tr key={order?.ID}>
              <td className="border px-4 py-2 text-center">{index+1}</td>
              <td className="border px-4 py-2 text-center">{order?.status}</td>
              <td className="border px-4 py-2 text-center">{order?.total}</td>
              <td className="border px-4 py-2 text-center">
                {order?.currencyUnit}
              </td>
              <td className="border px-4 py-2 flex justify-center items-center">
                <button className="bg-gray-600 hover:bg-gray-500 text-white font-bold py-2 px-4 rounded">
                  View
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Orders;
