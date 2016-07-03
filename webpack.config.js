module.exports = {
    entry: './ui/src/index.js',
    output: {
        filename: './ui/dist/app.js'
    },
    devtool: 'source-map',
    module: {
        loaders: [
            {
                test: /\.js?$/,
                exclude: /node_modules/,
                loader: 'babel',
                query: {
                    presets: ['react', 'es2015']
                }
            }
        ]
    },
    resolve: {
        extensions: ['', '.js', '.jsx']
    }
};
