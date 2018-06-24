import plotly
import plotly.graph_objs as go
import sys

if len(sys.argv) != 3:
    print('Please specify Plotly credentials')
    sys.exit(0)

plotly.tools.set_credentials_file(username=sys.argv[1], api_key=sys.argv[2])

files = {
    'RestClient_Proxy': 'rest_client/proxy.txt',
    'RestClient_Pure': 'rest_client/pure.txt',
    'RpcClient': 'rpc_client/rpc.txt'
}

data_small = []
data_big = []

for k, v in files.items():
    with open(v) as f:
        x = []
        y = []
        for line in f:
            (key, val) = line.split()
            func = key.split('_')[1]
            if func[-2] == '-':
                func = func[:-2]
            x.append(func)
            y.append(int(val))
        data_small.append(go.Bar(x=x[:5], y=y[:5], name=k))
        data_big.append(go.Bar(x=x[5:-1], y=y[5:-1], name=k))

layout = go.Layout(
    barmode='group'
)

fig = go.Figure(data=data_small, layout=layout)
plotly.plotly.plot(fig, filename='grpc_vs_rest_benchmark_small')


fig1 = go.Figure(data=data_big, layout=layout)
plotly.plotly.plot(fig1, filename='grpc_vs_rest_benchmark_big')
