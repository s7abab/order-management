import axios from "axios";

// get orders
export const getOrders = async () => {
  try {
    const res = await axios.get("http://localhost:8080/api/v1/order?id=11");
    console.log(res.data);
  } catch (error) {
    console.log(error);
  }
};
