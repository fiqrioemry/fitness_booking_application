import {
  FaGithub,
  FaLinkedin,
  FaTwitter,
  FaInstagram,
  FaEnvelope,
} from "react-icons/fa";
import { WebLogo } from "@/components/ui/WebLogo";

const Footer = () => {
  return (
    <footer className="bg-background text-muted-foreground border-t border-border">
      <div className="max-w-7xl mx-auto px-4 py-10 grid gap-10 grid-cols-1 md:grid-cols-4">
        {/* Brand */}
        <div>
          <WebLogo />
          <p className="text-sm mt-2 leading-relaxed">
            Your go-to space for high-energy classes and flexible fitness plans
            designed to match your lifestyle.
          </p>
        </div>

        {/* Navigation */}
        <div>
          <h3 className="text-base font-semibold mb-2 text-foreground">
            Quick Links
          </h3>
          <ul className="space-y-1 text-sm">
            {["Classes", "Packages", "Schedules", "About"].map((item) => (
              <li key={item}>
                <a
                  href={`/${item.toLowerCase()}`}
                  className="hover:text-primary transition-colors"
                >
                  {item}
                </a>
              </li>
            ))}
          </ul>
        </div>

        {/* Newsletter */}
        <div>
          <h3 className="text-base font-semibold mb-2 text-foreground">
            Join Our Community
          </h3>
          <p className="text-sm mb-3">
            Get tips, exclusive offers, and new class alerts directly to your
            inbox.
          </p>
          <form className="flex flex-col gap-2">
            <input
              type="email"
              placeholder="Enter your email"
              className="input"
            />
            <button
              type="submit"
              className="bg-primary text-primary-foreground text-sm font-medium px-4 py-2 rounded-md hover:bg-primary/90 transition"
            >
              Subscribe
            </button>
          </form>
        </div>

        {/* Social Media */}
        <div>
          <h3 className="text-base font-semibold mb-2 text-foreground">
            Connect with Us
          </h3>
          <div className="flex gap-4 mt-3">
            {[
              {
                href: "https://instagram.com/yourstudio",
                icon: <FaInstagram />,
              },
              {
                href: "https://linkedin.com/in/yourstudio",
                icon: <FaLinkedin />,
              },
              {
                href: "https://twitter.com/yourstudio",
                icon: <FaTwitter />,
              },
              {
                href: "https://github.com/yourstudio",
                icon: <FaGithub />,
              },
              {
                href: "mailto:info@fitbookstudio.com",
                icon: <FaEnvelope />,
              },
            ].map(({ href, icon }, i) => (
              <a
                key={i}
                href={href}
                target="_blank"
                rel="noreferrer"
                className="text-xl text-muted-foreground hover:text-primary transition-colors"
              >
                {icon}
              </a>
            ))}
          </div>
        </div>
      </div>

      <div className="py-6 border-t border-border text-center text-xs text-muted-foreground">
        &copy; 2025 FitBook Studio. All rights reserved.
      </div>
    </footer>
  );
};

export default Footer;
