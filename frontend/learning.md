## Install webpack

```
npm inits
npm install -d webpack webpack-cli webpack-dev-server
npm install -D @webpack-cli/generators
```

## Init project

```
npx webpack-cli init
```

? Which of the following JS solutions do you want to use? Typescript
? Do you want to use webpack-dev-server? Yes
? Do you want to simplify the creation of HTML files for your bundle? Yes
? Do you want to add PWA support? No
? Which of the following CSS solutions do you want to use? CSS only
? Will you be using PostCSS in your project? No
? Do you want to extract CSS for every file? Yes
? Do you like to install prettier to format generated configuration? Yes
? Pick a package manager: yarn

# Run dev server

```
npx webpack server --open --static-directory dist --mode=development

```

## Install Typescript

```
npm install -d typescript ts-loader
npm install react react-dom @types/react @types/react-dom
npx tsc --init
```

## Install React

```
npm install react react-dom @types/react @types/react-dom
```

## Install Babel

```
npm install -d babel-loader @babel/core @babel/preset-env @babel/preset-react
```

## Init eslint

```
npx eslint --init
```

## Edit webpack.config

```js
const config = {
  plugins: [
    new HtmlWebpackPlugin({
      template: './public/index.html',
      filename: 'index.html',
    }),
  ],
};
```
