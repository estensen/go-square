package square_test

import (
	"bytes"
	"testing"

	"github.com/celestiaorg/go-square/pkg/shares"
	"github.com/celestiaorg/go-square/pkg/square"
	"github.com/stretchr/testify/require"
)

// func TestBuilderSquareSizeEstimation(t *testing.T) {
// 	type test struct {
// 		name               string
// 		normalTxs          int
// 		pfbCount, pfbSize  int
// 		expectedSquareSize int
// 	}
// 	tests := []test{
// 		{"empty block", 0, 0, 0, appconsts.MinSquareSize},
// 		{"one normal tx", 1, 0, 0, 1},
// 		{"one small pfb small block", 0, 1, 100, 2},
// 		{"mixed small block", 10, 12, 500, 8},
// 		{"small block 2", 0, 12, 1000, 8},
// 		{"mixed medium block 2", 10, 20, 10000, 32},
// 		{"one large pfb large block", 0, 1, 1000000, 64},
// 		{"one hundred large pfb large block", 0, 100, 100000, appconsts.DefaultGovMaxSquareSize},
// 		{"one hundred large pfb medium block", 100, 100, 100000, appconsts.DefaultGovMaxSquareSize},
// 		{"mixed transactions large block", 100, 100, 100000, appconsts.DefaultGovMaxSquareSize},
// 		{"mixed transactions large block 2", 1000, 1000, 10000, appconsts.DefaultGovMaxSquareSize},
// 		{"mostly transactions large block", 10000, 1000, 100, appconsts.DefaultGovMaxSquareSize},
// 		{"only small pfb large block", 0, 10000, 1, appconsts.DefaultGovMaxSquareSize},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			rand := tmrand.NewRand()
// 			signer, err := testnode.NewOfflineSigner()
// 			require.NoError(t, err)
// 			txs := generateMixedTxs(signer, rand, tt.normalTxs, tt.pfbCount, 1, tt.pfbSize)
// 			square, _, err := square.Build(txs, appconsts.LatestVersion, appconsts.DefaultGovMaxSquareSize)
// 			require.NoError(t, err)
// 			require.EqualValues(t, tt.expectedSquareSize, square.Size())
// 		})
// 	}
// }

func TestBuilderRejectsTransactions(t *testing.T) {
	builder, err := square.NewBuilder(2, 64) // 2 x 2 square
	require.NoError(t, err)
	require.False(t, builder.AppendTx(newTx(shares.AvailableBytesFromCompactShares(4)+1)))
	require.True(t, builder.AppendTx(newTx(shares.AvailableBytesFromCompactShares(4))))
	require.False(t, builder.AppendTx(newTx(1)))
}

// func TestBuilderRejectsBlobTransactions(t *testing.T) {
// 	ns1 := namespace.MustNewV0(bytes.Repeat([]byte{1}, namespace.NamespaceVersionZeroIDSize))
// 	testCases := []struct {
// 		blobSize []int
// 		added    bool
// 	}{
// 		{
// 			blobSize: []int{shares.AvailableBytesFromSparseShares(3) + 1},
// 			added:    false,
// 		},
// 		{
// 			blobSize: []int{shares.AvailableBytesFromSparseShares(3)},
// 			added:    true,
// 		},
// 		{
// 			blobSize: []int{shares.AvailableBytesFromSparseShares(2) + 1, shares.AvailableBytesFromSparseShares(1)},
// 			added:    false,
// 		},
// 		{
// 			blobSize: []int{shares.AvailableBytesFromSparseShares(1), shares.AvailableBytesFromSparseShares(1)},
// 			added:    true,
// 		},
// 		{
// 			// fun fact: three blobs increases the size of the PFB to two shares, hence this fails
// 			blobSize: []int{
// 				shares.AvailableBytesFromSparseShares(1),
// 				shares.AvailableBytesFromSparseShares(1),
// 				shares.AvailableBytesFromSparseShares(1),
// 			},
// 			added: false,
// 		},
// 	}

// 	for idx, tc := range testCases {
// 		t.Run(fmt.Sprintf("case%d", idx), func(t *testing.T) {
// 			builder, err := square.NewBuilder(2, 64)
// 			require.NoError(t, err)
// 			txs := generateBlobTxsWithNamespaces(t, ns1.Repeat(len(tc.blobSize)), [][]int{tc.blobSize})
// 			require.Len(t, txs, 1)
// 			blobTx, isBlobTx := blob.UnmarshalBlobTx(txs[0])
// 			require.True(t, isBlobTx)
// 			require.Equal(t, tc.added, builder.AppendBlobTx(blobTx))
// 		})
// 	}
// }

func TestBuilderInvalidConstructor(t *testing.T) {
	_, err := square.NewBuilder(-4, 64)
	require.Error(t, err)
	_, err = square.NewBuilder(0, 64)
	require.Error(t, err)
	_, err = square.NewBuilder(13, 64)
	require.Error(t, err)
}

func newTx(len int) []byte {
	return bytes.Repeat([]byte{0}, shares.RawTxSize(len))
}

// func TestBuilderFindTxShareRange(t *testing.T) {
// 	signer, err := testnode.NewOfflineSigner()
// 	require.NoError(t, err)
// 	blockTxs := testfactory.GenerateRandomTxs(5, 900).ToSliceOfBytes()
// 	blockTxs = append(blockTxs, blobfactory.RandBlobTxsRandomlySized(signer, tmrand.NewRand(), 5, 1000, 10).ToSliceOfBytes()...)
// 	require.Len(t, blockTxs, 10)

// 	builder, err := square.NewBuilder(128, 64, blockTxs...)
// 	require.NoError(t, err)

// 	dataSquare, err := builder.Export()
// 	require.NoError(t, err)
// 	size := dataSquare.Size() * dataSquare.Size()

// 	var lastEnd int
// 	for idx, tx := range blockTxs {
// 		blobTx, isBlobTx := blob.UnmarshalBlobTx(tx)
// 		if isBlobTx {
// 			tx = blobTx.Tx
// 		}
// 		shareRange, err := builder.FindTxShareRange(idx)
// 		require.NoError(t, err)
// 		if idx == 5 {
// 			// normal txs and PFBs use a different namespace so there
// 			// can't be any overlap in the index
// 			require.Greater(t, shareRange.Start, lastEnd-1)
// 		} else {
// 			require.GreaterOrEqual(t, shareRange.Start, lastEnd-1)
// 		}
// 		require.LessOrEqual(t, shareRange.End, size)
// 		txShares := dataSquare[shareRange.Start : shareRange.End+1]
// 		parsedShares, err := rawData(txShares)
// 		require.NoError(t, err)
// 		require.True(t, bytes.Contains(parsedShares, tx))
// 		lastEnd = shareRange.End
// 	}
// }

func rawData(shares []shares.Share) ([]byte, error) {
	var data []byte
	for _, share := range shares {
		rawData, err := share.RawData()
		if err != nil {
			return nil, err
		}
		data = append(data, rawData...)
	}
	return data, nil
}

// // TestSquareBlobPositions ensures that the share commitment rules which dictate the padding
// // between blobs is followed as well as the ordering of blobs by namespace.
// func TestSquareBlobPostions(t *testing.T) {
// 	ns1 := namespace.MustNewV0(bytes.Repeat([]byte{1}, namespace.NamespaceVersionZeroIDSize))
// 	ns2 := namespace.MustNewV0(bytes.Repeat([]byte{2}, namespace.NamespaceVersionZeroIDSize))
// 	ns3 := namespace.MustNewV0(bytes.Repeat([]byte{3}, namespace.NamespaceVersionZeroIDSize))

// 	type test struct {
// 		squareSize      int
// 		blobTxs         [][]byte
// 		expectedIndexes [][]uint32
// 	}
// 	tests := []test{
// 		{
// 			squareSize: 4,
// 			blobTxs: generateBlobTxsWithNamespaces(
// 				t,
// 				[]namespace.Namespace{ns1},
// 				[][]int{{1}},
// 			),
// 			expectedIndexes: [][]uint32{{1}},
// 		},
// 		{
// 			squareSize: 4,
// 			blobTxs: generateBlobTxsWithNamespaces(
// 				t,
// 				[]namespace.Namespace{ns1, ns1},
// 				repeat([]int{100}, 2),
// 			),
// 			expectedIndexes: [][]uint32{{2}, {3}},
// 		},
// 		{
// 			squareSize: 4,
// 			blobTxs: generateBlobTxsWithNamespaces(
// 				t,
// 				[]namespace.Namespace{ns1, ns1, ns1, ns1, ns1, ns1, ns1, ns1, ns1},
// 				repeat([]int{100}, 9),
// 			),
// 			expectedIndexes: [][]uint32{{7}, {8}, {9}, {10}, {11}, {12}, {13}, {14}, {15}},
// 		},
// 		{
// 			squareSize: 4,
// 			blobTxs: generateBlobTxsWithNamespaces(
// 				t,
// 				[]namespace.Namespace{ns1, ns1, ns1},
// 				[][]int{{10000}, {10000}, {1000000}},
// 			),
// 			expectedIndexes: [][]uint32{},
// 		},
// 		{
// 			squareSize: 64,
// 			blobTxs: generateBlobTxsWithNamespaces(
// 				t,
// 				[]namespace.Namespace{ns1, ns1, ns1},
// 				[][]int{{1000}, {10000}, {10000}},
// 			),
// 			expectedIndexes: [][]uint32{{3}, {6}, {27}},
// 		},
// 		{
// 			squareSize: 32,
// 			blobTxs: generateBlobTxsWithNamespaces(
// 				t,
// 				[]namespace.Namespace{ns2, ns1, ns1},
// 				[][]int{{100}, {100}, {100}},
// 			),
// 			expectedIndexes: [][]uint32{{5}, {3}, {4}},
// 		},
// 		{
// 			squareSize: 16,
// 			blobTxs: generateBlobTxsWithNamespaces(
// 				t,
// 				[]namespace.Namespace{ns1, ns2, ns1},
// 				[][]int{{100}, {900}, {900}}, // 1, 2, 2 shares respectively
// 			),
// 			expectedIndexes: [][]uint32{{3}, {6}, {4}},
// 		},
// 		{
// 			squareSize: 4,
// 			blobTxs: generateBlobTxsWithNamespaces(
// 				t,
// 				[]namespace.Namespace{ns1, ns3, ns3, ns2},
// 				[][]int{{100}, {1000, 1000}, {420}},
// 			),
// 			expectedIndexes: [][]uint32{{3}, {5, 8}, {4}},
// 		},
// 		{
// 			// no blob txs should make it in the square
// 			squareSize: 1,
// 			blobTxs: generateBlobTxsWithNamespaces(
// 				t,
// 				[]namespace.Namespace{ns1, ns2, ns3},
// 				[][]int{{1000}, {1000}, {1000}},
// 			),
// 			expectedIndexes: [][]uint32{},
// 		},
// 		{
// 			// only two blob txs should make it in the square (after reordering)
// 			squareSize: 4,
// 			blobTxs: generateBlobTxsWithNamespaces(
// 				t,
// 				[]namespace.Namespace{ns3, ns2, ns1},
// 				[][]int{{2000}, {2000}, {5000}},
// 			),
// 			expectedIndexes: [][]uint32{{7}, {2}},
// 		},
// 		{
// 			squareSize: 4,
// 			blobTxs: generateBlobTxsWithNamespaces(
// 				t,
// 				[]namespace.Namespace{ns3, ns3, ns2, ns1},
// 				[][]int{{1800, 1000}, {22000}, {1800}},
// 			),
// 			// should be ns1 and {ns3, ns3} as ns2 is too large
// 			expectedIndexes: [][]uint32{{6, 10}, {2}},
// 		},
// 		{
// 			squareSize: 4,
// 			blobTxs: generateBlobTxsWithNamespaces(
// 				t,
// 				[]namespace.Namespace{ns1, ns3, ns3, ns1, ns2, ns2},
// 				[][]int{{100}, {1400, 900, 200, 200}, {420}},
// 			),
// 			expectedIndexes: [][]uint32{{3}, {7, 10, 4, 5}, {6}},
// 		},
// 		{
// 			squareSize: 4,
// 			blobTxs: generateBlobTxsWithNamespaces(
// 				t,
// 				[]namespace.Namespace{ns1, ns3, ns3, ns1, ns2, ns2},
// 				[][]int{{100}, {900, 1400, 200, 200}, {420}},
// 			),
// 			expectedIndexes: [][]uint32{{3}, {7, 9, 4, 5}, {6}},
// 		},
// 		{
// 			squareSize: 16,
// 			blobTxs: generateBlobTxsWithNamespaces(
// 				t,
// 				[]namespace.Namespace{ns1, ns1},
// 				[][]int{{100}, {shares.AvailableBytesFromSparseShares(64)}},
// 			),
// 			// There should be one share padding between the two blobs
// 			expectedIndexes: [][]uint32{{2}, {3}},
// 		},
// 		{
// 			squareSize: 16,
// 			blobTxs: generateBlobTxsWithNamespaces(
// 				t,
// 				[]namespace.Namespace{ns1, ns1},
// 				[][]int{{100}, {shares.AvailableBytesFromSparseShares(64) + 1}},
// 			),
// 			// There should be one share padding between the two blobs
// 			expectedIndexes: [][]uint32{{2}, {4}},
// 		},
// 	}
// 	for i, tt := range tests {
// 		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
// 			builder, err := square.NewBuilder(tt.squareSize, 64)
// 			require.NoError(t, err)
// 			for _, tx := range tt.blobTxs {
// 				blobTx, isBlobTx := blob.UnmarshalBlobTx(tx)
// 				require.True(t, isBlobTx)
// 				_ = builder.AppendBlobTx(blobTx)
// 			}
// 			square, err := builder.Export()
// 			require.NoError(t, err)
// 			txs, err := shares.ParseTxs(square)
// 			require.NoError(t, err)
// 			for j, tx := range txs {
// 				wrappedPFB, isWrappedPFB := blob.UnmarshalIndexWrapper(tx)
// 				require.True(t, isWrappedPFB)
// 				require.Equal(t, tt.expectedIndexes[j], wrappedPFB.ShareIndexes, j)
// 			}
// 		})
// 	}
// }

// func repeat[T any](s T, count int) []T {
// 	ss := make([]T, count)
// 	for i := 0; i < count; i++ {
// 		ss[i] = s
// 	}
// 	return ss
// }
