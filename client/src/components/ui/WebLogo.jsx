import { Link } from "react-router-dom";

const WebLogo = () => {
  return (
    <Link to="/">
      <img src="/logo.png" className="h-12" alt="logo" />
    </Link>
  );
};

export { WebLogo };
