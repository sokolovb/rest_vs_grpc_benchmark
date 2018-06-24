import plotly
import plotly.graph_objs as go
import sys

if len(sys.argv) != 3:
    print('Please specify Plotly credentials')
    sys.exit(0)

plotly.tools.set_credentials_file(username=sys.argv[1], api_key=sys.argv[2])

data_serialize_small = []
data_serialize_big = []
data_deserialize_small = []
data_deserialize_big = []

with open('serialization/serialization.txt') as f:
    x = []
    y = []
    for line in f:
        (key, val) = line.split()
        x.append(key.split('_')[1])
        y.append(int(val))
    data_serialize_small.append(go.Bar(x=x[:5], y=y[:5], name='BenchmarkSerializeJSON'))
    data_serialize_small.append(go.Bar(x=x[8:13], y=y[8:13], name='BenchmarkSerializeProtobuf'))
    data_serialize_big.append(go.Bar(x=x[5:8], y=y[5:8], name='BenchmarkSerializeJSON'))
    data_serialize_big.append(go.Bar(x=x[13:16], y=y[13:16], name='BenchmarkSerializeProtobuf'))
    data_deserialize_small.append(go.Bar(x=x[16:21], y=y[16:21], name='BenchmarkDeserializeJSON'))
    data_deserialize_small.append(go.Bar(x=x[24:29], y=y[24:29], name='BenchmarkDeserializeProtobuf'))
    data_deserialize_big.append(go.Bar(x=x[21:24], y=y[21:24], name='BenchmarkDeserializeJSON'))
    data_deserialize_big.append(go.Bar(x=x[29:], y=y[29:], name='BenchmarkDeserializeProtobuf'))

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
