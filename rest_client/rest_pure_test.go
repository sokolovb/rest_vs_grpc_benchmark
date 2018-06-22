package rest_client_test

import "testing"

// pure rest api benchmarks

func BenchmarkRestClient_GetInt_Pure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restPure.GetInt(); err != nil {
			b.Fatalf("GetInt() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetIntSlice_Pure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restPure.GetIntSlice(); err != nil {
			b.Fatalf("GetIntSlice() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetString_Pure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restPure.GetString(); err != nil {
			b.Fatalf("GetString() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetStringSlice_Pure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restPure.GetStringSlice(); err != nil {
			b.Fatalf("GetStringSlice() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetStruct_Pure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restPure.GetStruct(); err != nil {
			b.Fatalf("GetStruct() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetBlob_Pure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restPure.GetBlob(); err != nil {
			b.Fatalf("GetBlob() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetStructSlices_Pure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restPure.GetStructSlices(); err != nil {
			b.Fatalf("GetStructSlices() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetStructStructs_Pure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restPure.GetStructStructs(); err != nil {
			b.Fatalf("GetStructStructs() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetFile_Pure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restPure.GetFile(); err != nil {
			b.Fatalf("GetFile() error: %v", err)
		}
	}
}
