package square_test

// func BenchmarkSquareConstruct(b *testing.B) {
// 	for _, txCount := range []int{10, 100, 1000} {
// 		b.Run(fmt.Sprintf("txCount=%d", txCount), func(b *testing.B) {
// 			signer, err := testnode.NewOfflineSigner()
// 			require.NoError(b, err)
// 			txs := generateOrderedTxs(signer, tmrand.NewRand(), txCount/2, txCount/2, 1, 1024)
// 			b.ResetTimer()
// 			for i := 0; i < b.N; i++ {
// 				_, err := square.Construct(txs, LatestVersion, DefaultSquareSizeUpperBound)
// 				require.NoError(b, err)
// 			}
// 		})
// 	}
// }

// func BenchmarkSquareBuild(b *testing.B) {
// 	for _, txCount := range []int{10, 100, 1000, 10000} {
// 		b.Run(fmt.Sprintf("txCount=%d", txCount), func(b *testing.B) {
// 			signer, err := testnode.NewOfflineSigner()
// 			require.NoError(b, err)
// 			txs := generateMixedTxs(signer, tmrand.NewRand(), txCount/2, txCount/2, 1, 1024)
// 			b.ResetTimer()
// 			for i := 0; i < b.N; i++ {
// 				_, _, err := square.Build(txs, LatestVersion, DefaultSquareSizeUpperBound)
// 				require.NoError(b, err)
// 			}
// 		})
// 	}
// 	const txCount = 10
// 	for _, blobSize := range []int{10, 100, 1000, 10000} {
// 		b.Run(fmt.Sprintf("blobSize=%d", blobSize), func(b *testing.B) {
// 			signer, err := testnode.NewOfflineSigner()
// 			require.NoError(b, err)
// 			txs := generateMixedTxs(signer, tmrand.NewRand(), 0, txCount, 1, blobSize)
// 			b.ResetTimer()
// 			for i := 0; i < b.N; i++ {
// 				_, _, err := square.Build(txs, LatestVersion, DefaultSquareSizeUpperBound)
// 				require.NoError(b, err)
// 			}
// 		})
// 	}
// }
