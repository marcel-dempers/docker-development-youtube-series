// using System.Diagnostics;

// /// <summary>
// /// It is recommended to use a custom type to hold references for ActivitySource.
// /// This avoids possible type collisions with other components in the DI container.
// /// </summary>
// public class Instrumentation : IDisposable
// {
//     internal const string ActivitySourceName = "videos-api";
//     internal const string ActivitySourceVersion = "1.0.0";

//     public Instrumentation()
//     {
//         this.ActivitySource = new ActivitySource(ActivitySourceName, ActivitySourceVersion);
//     }

//     public ActivitySource ActivitySource { get; }

//     public void Dispose()
//     {
//         this.ActivitySource.Dispose();
//     }
// }
