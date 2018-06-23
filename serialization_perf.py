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

data_serialize_small = []
data_serialize_big = []
data_deserialize_small = []
data_deserialize_big = []

with open('serialization/serialization.txt') as f:
    x = []
    y = []
    for line in f:
        (key, val) = line.split()
        x.append(key[9:])
        y.append(int(val))
    data_serialize_small.append(go.Bar(x=x[:5], y=y[:5], name='BenchmarkSerializeJSON'))
    data_serialize_small.append(go.Bar(x=x[8:12], y=y[8:12], name='BenchmarkSerializeProtobuf'))
    data_serialize_big.append(go.Bar(x=x[5:8], y=y[5:8], name='BenchmarkSerializeJSON'))
    data_serialize_big.append(go.Bar(x=x[12:15], y=y[12:15], name='BenchmarkSerializeProtobuf'))
    data_deserialize_small.append(go.Bar(x=x[15:20], y=y[15:20], name='BenchmarkDeserializeJSON'))
    data_deserialize_small.append(go.Bar(x=x[23:28], y=y[23:28], name='BenchmarkDeserializeProtobuf'))
    data_deserialize_big.append(go.Bar(x=x[20:23], y=y[20:23], name='BenchmarkDeserializeJSON'))
    data_deserialize_big.append(go.Bar(x=x[28:], y=y[28:], name='BenchmarkDeserializeProtobuf'))

layout = go.Layout(
    barmode='group'
)

fig = go.Figure(data=data_serialize_small, layout=layout)
plotly.plotly.plot(fig, filename='serialize_small')


fig1 = go.Figure(data=data_serialize_big, layout=layout)
plotly.plotly.plot(fig1, filename='serialize_big')


fig2 = go.Figure(data=data_deserialize_small, layout=layout)
plotly.plotly.plot(fig2, filename='deserialize_small')


fig3 = go.Figure(data=data_deserialize_big, layout=layout)
plotly.plotly.plot(fig3, filename='deserialize_big')
