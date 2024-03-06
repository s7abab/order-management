import ProductCard from "./ProductCard";

const Products = ({ products, close }) => {
  return (
    <div className="absolute top-[50%] left-[50%] translate-x-[-50%] translate-y-[-50%] p-4 px-2 bg-white rounded-lg shadow-xl w-[500px]">
      <button className="font-semibold absolute right-4 top-1" onClick={close}>
        X
      </button>
      {products.map((product) => (
        <div className="m-2">
          <ProductCard product={product} />
        </div>
      ))}
    </div>
  );
};

export default Products;
