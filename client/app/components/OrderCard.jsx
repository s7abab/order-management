const OrderCard = ({ order }) => {
  return (
    <div className="bg-gray-200 rounded-lg shadow-md p-4">
      <div className="font-semibold text-lg mb-2">{order.status}</div>
      <div className="mt-4 flex justify-between items-center">
        <div className="text-gray-500 font-bold">{order.total}</div>
        <div className="text-gray-500 font-bold">{order.currencyUnit}</div>
      </div>
    </div>
  );
};

export default OrderCard;
