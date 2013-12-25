package main

import(
	// "fmt"
	"testing"
)

// func BenchmarkQueryAllSingleTag(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllSingleTag()
// 	}
// }

// func BenchmarkQueryAllSingleTagDistinct(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllSingleTagDistinct()
// 	}
// }

// func BenchmarkQueryAllSingleTagGrouped(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllSingleTagGrouped()
// 	}
// }

func BenchmarkQueryAllSingleTagSegmented(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllSingleTagSegmented()
	}
}

func BenchmarkQueryAllSingleTagSegmentedDistinct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllSingleTagSegmentedDistinct()
	}
}

func BenchmarkQueryAllSingleTagSegmentedGrouped(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllSingleTagSegmentedGrouped()
	}
}

// func BenchmarkQueryAllTwoTags(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllTwoTags()
// 	}
// }

// func BenchmarkQueryAllTwoTagsDistinct(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllTwoTagsDistinct()
// 	}
// }

// func BenchmarkQueryAllTwoTagsGrouped(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllTwoTagsGrouped()
// 	}
// }

func BenchmarkQueryAllTwoTagsSegmented(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTwoTagsSegmented()
	}
}

func BenchmarkQueryAllTwoTagsSegmentedDistinct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTwoTagsSegmentedDistinct()
	}
}

func BenchmarkQueryAllTwoTagsSegmentedGrouped(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTwoTagsSegmentedGrouped()
	}
}


// func BenchmarkQueryAllThreeTags(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllThreeTags()
// 	}
// }

// func BenchmarkQueryAllThreeTagsDistinct(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllThreeTagsDistinct()
// 	}
// }

// func BenchmarkQueryAllThreeTagsGrouped(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllThreeTagsGrouped()
// 	}
// }

func BenchmarkQueryAllThreeTagsSegmented(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllThreeTagsSegmented()
	}
}

func BenchmarkQueryAllThreeTagsSegmentedDistinct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllThreeTagsSegmentedDistinct()
	}
}

func BenchmarkQueryAllThreeTagsSegmentedGrouped(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllThreeTagsSegmentedGrouped()
	}
}

func BenchmarkQueryAllTenTagsSegmented(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsSegmented()
	}
}

func BenchmarkQueryAllTenTagsSegmentedDistinct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsSegmentedDistinct()
	}
}

func BenchmarkQueryAllTenTagsSegmentedGrouped(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsSegmentedGrouped()
	}
}

func BenchmarkQueryAllTenTagsSegmentedGroupedParallel2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTagsSegmentedGroupedParallelX(2)
	}
}

func BenchmarkQueryAllTenTagsSegmentedGroupedParallel10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTagsSegmentedGroupedParallelX(10)
	}
}

func BenchmarkQueryAllTenTagsSegmentedGroupedParallel20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTagsSegmentedGroupedParallelX(20)
	}
}