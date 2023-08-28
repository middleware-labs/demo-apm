const path = require('path');
module.exports = {
    entry: './src/index.ts',
    target: 'webworker',
    mode: "development",
    module: {
        rules: [
            {
                test: /\.ts$/,
                use: 'ts-loader',
                exclude: /node_modules/,
            },
            {
                test: /\.m?js/,
                resolve: {
                    fullySpecified: false
                }
            }
        ],
    },
    resolve: {
        extensions: ['.ts', '.js'],
    },
    output: {
        filename: 'worker.js',
        path: path.resolve(__dirname, 'dist'),
    },
};
