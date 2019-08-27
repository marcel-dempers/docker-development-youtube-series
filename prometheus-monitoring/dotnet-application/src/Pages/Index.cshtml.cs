using System;
using System.Diagnostics;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using Prometheus;
namespace work.Pages
{
    public class IndexModel : PageModel
    {
        private static readonly Counter ProcessedJobCount = Metrics
	.CreateCounter("dotnet_request_operations_total", "The total number of processed requests");
        public void OnGet()
        {
            var sw = Stopwatch.StartNew();

            sw.Stop();
            ProcessedJobCount.Inc();
            var histogram =
            Metrics.CreateHistogram(
                    "dotnet_request_duration_seconds",
                    "Histogram for the duration in seconds.",
                    new HistogramConfiguration
                    {
                        Buckets = Histogram.LinearBuckets(start: 1, width: 1, count: 5)
                    });

            histogram.Observe(sw.Elapsed.TotalSeconds);

        }
    }
}
