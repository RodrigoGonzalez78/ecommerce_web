-- Estructura de tabla para la tabla `address`
CREATE TABLE address (
  id SERIAL PRIMARY KEY,
  street TEXT NOT NULL,
  postal_code INT NOT NULL,
  neighborhood TEXT NOT NULL,
  city TEXT NOT NULL
);

-- Estructura de tabla para la tabla `categories`
CREATE TABLE categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL
);

-- Estructura de tabla para la tabla `consults`
CREATE TABLE consults (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  email VARCHAR(50) NOT NULL,
  description VARCHAR(150) NOT NULL,
  attended VARCHAR(2) NOT NULL DEFAULT 'NO',
  archived TEXT NOT NULL DEFAULT 'NO'
);

-- Estructura de tabla para la tabla `products`
CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  price NUMERIC(15, 2) NOT NULL,
  stock INT NOT NULL,
  description VARCHAR(600) NOT NULL,
  down VARCHAR(2) NOT NULL DEFAULT 'NO',
  image TEXT NOT NULL,
  id_categorie INT NOT NULL,
  FOREIGN KEY (id_categorie) REFERENCES categories(id)
);

-- Estructura de tabla para la tabla `profile`
CREATE TABLE profile (
  id SERIAL PRIMARY KEY,
  descrition VARCHAR(100) NOT NULL
);

-- Estructura de tabla para la tabla `sales`
CREATE TABLE sales (
  id SERIAL PRIMARY KEY,
  id_user INT NOT NULL,
  total_price NUMERIC(15, 2) NOT NULL,
  date DATE NOT NULL DEFAULT CURRENT_DATE,
  FOREIGN KEY (id_user) REFERENCES users(id)
);

-- Estructura de tabla para la tabla `salesdetails`
CREATE TABLE sale_details (
  id SERIAL PRIMARY KEY,
  id_sale INT NOT NULL,
  count INT NOT NULL,
  price NUMERIC(15, 2) NOT NULL,
  id_product INT NOT NULL,
  FOREIGN KEY (id_sale) REFERENCES sales(id),
  FOREIGN KEY (id_product) REFERENCES products(id)
);

-- Estructura de tabla para la tabla `users`
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  email VARCHAR(100) NOT NULL,
  password VARCHAR(200) NOT NULL,
  down TEXT NOT NULL DEFAULT 'NO',
  id_address INT,
  id_profile INT NOT NULL DEFAULT 2,
  FOREIGN KEY (id_profile) REFERENCES profile(id),
  FOREIGN KEY (id_address) REFERENCES address(id)
);
