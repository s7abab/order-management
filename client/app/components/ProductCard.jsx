const ProductCard = ({ product }) => {
  return (
    <div className="bg-gray-200 rounded-lg shadow-md p-4">
      <div className="font-semibold text-lg mb-2">Name: {product.name}</div>
      <div className="text-gray-500">Description: {product.description}</div>
      <div className="mt-4 flex justify-between items-center">
        <div className="text-gray-500 font-bold">Price : {product.price}</div>
        <div className="text-gray-500 font-bold">Qty : {product.quantity}</div>
      </div>
    </div>
  );
};

export default ProductCard;
