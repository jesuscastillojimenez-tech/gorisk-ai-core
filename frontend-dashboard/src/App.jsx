import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { ShieldCheck, Brain, Server, RefreshCw, Send, AlertCircle, CheckCircle2, XCircle } from 'lucide-react';

const API_BASE_URL = 'http://localhost:8080/api/v1/applications';

export default function App() {
  const [applications, setApplications] = useState([]);
  const [loading, setLoading] = useState(false);
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState(null);

  const [formData, setFormData] = useState({
    applicant_name: '',
    monthly_income: '',
    requested_amount: ''
  });

  const fetchApplications = async () => {
    setLoading(true);
    try {
      const response = await axios.get(API_BASE_URL);
      setApplications(response.data.data || response.data || []);
      setError(null);
    } catch (err) {
      setError('Could not connect to Go Backend on port 8080.');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchApplications();
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setSubmitting(true);
    setError(null);

    const payload = {
      applicant_name: formData.applicant_name,
      monthly_income: parseFloat(formData.monthly_income),
      requested_amount: parseFloat(formData.requested_amount)
    };

    try {
      await axios.post(API_BASE_URL, payload);
      setFormData({ applicant_name: '', monthly_income: '', requested_amount: '' });
      await fetchApplications();
    } catch (err) {
      setError('Failed to submit application to AI engine.');
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div className="min-h-screen bg-slate-950 text-slate-100 font-sans p-6 md:p-10">
      <header className="max-w-7xl mx-auto flex flex-col md:flex-row justify-between items-start md:items-center mb-10 pb-6 border-b border-slate-800 gap-4">
        <div>
          <div className="flex items-center gap-3">
            <ShieldCheck className="w-8 h-8 text-emerald-400" />
            <h1 className="text-3xl font-bold tracking-tight bg-gradient-to-r from-emerald-400 to-cyan-400 bg-clip-text text-transparent">
              GoRisk AI Core
            </h1>
          </div>
          <p className="text-slate-400 text-sm mt-1">Enterprise Risk Evaluation & Event-Driven Pipeline Dashboard</p>
        </div>

        <div className="flex items-center gap-4 text-xs font-mono">
          <div className="flex items-center gap-2 px-3 py-1.5 rounded-full bg-slate-900 border border-slate-800">
            <span className="w-2 h-2 rounded-full bg-emerald-500 animate-pulse"></span>
            <Server className="w-3.5 h-3.5 text-slate-400" /> Go Service: :8080
          </div>
          <div className="flex items-center gap-2 px-3 py-1.5 rounded-full bg-slate-900 border border-slate-800">
            <Brain className="w-3.5 h-3.5 text-purple-400" /> Gemini AI Engine
          </div>
        </div>
      </header>

      <main className="max-w-7xl mx-auto grid grid-cols-1 lg:grid-cols-3 gap-8">
        <section className="bg-slate-900 p-6 rounded-2xl border border-slate-800 shadow-xl h-fit">
          <h2 className="text-lg font-semibold mb-4 flex items-center gap-2 text-slate-200">
            <Send className="w-5 h-5 text-emerald-400" /> New Credit Request
          </h2>

          {error && (
            <div className="mb-4 p-3 rounded-lg bg-rose-950/50 border border-rose-800 text-rose-300 text-xs flex items-center gap-2">
              <AlertCircle className="w-4 h-4 shrink-0" />
              <span>{error}</span>
            </div>
          )}

          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <label className="block text-xs font-medium text-slate-400 mb-1">Applicant Name</label>
              <input
                type="text"
                required
                placeholder="e.g. Jesus Castillo"
                value={formData.applicant_name}
                onChange={(e) => setFormData({ ...formData, applicant_name: e.target.value })}
                className="w-full bg-slate-950 border border-slate-800 rounded-lg px-3 py-2 text-sm text-slate-200 focus:outline-none focus:border-emerald-500 transition-colors"
              />
            </div>

            <div>
              <label className="block text-xs font-medium text-slate-400 mb-1">Monthly Income ($)</label>
              <input
                type="number"
                required
                placeholder="e.g. 50000"
                value={formData.monthly_income}
                onChange={(e) => setFormData({ ...formData, monthly_income: e.target.value })}
                className="w-full bg-slate-950 border border-slate-800 rounded-lg px-3 py-2 text-sm text-slate-200 focus:outline-none focus:border-emerald-500 transition-colors"
              />
            </div>

            <div>
              <label className="block text-xs font-medium text-slate-400 mb-1">Requested Amount ($)</label>
              <input
                type="number"
                required
                placeholder="e.g. 150000"
                value={formData.requested_amount}
                onChange={(e) => setFormData({ ...formData, requested_amount: e.target.value })}
                className="w-full bg-slate-950 border border-slate-800 rounded-lg px-3 py-2 text-sm text-slate-200 focus:outline-none focus:border-emerald-500 transition-colors"
              />
            </div>

            <button
              type="submit"
              disabled={submitting}
              className="w-full mt-2 bg-emerald-500 hover:bg-emerald-600 disabled:bg-slate-800 text-slate-950 font-semibold py-2.5 px-4 rounded-lg text-sm transition-all flex items-center justify-center gap-2 shadow-lg shadow-emerald-500/10"
            >
              {submitting ? (
                <>
                  <RefreshCw className="w-4 h-4 animate-spin text-slate-950" />
                  Evaluating with AI...
                </>
              ) : (
                'Submit Application'
              )}
            </button>
          </form>
        </section>

        <section className="lg:col-span-2 bg-slate-900 p-6 rounded-2xl border border-slate-800 shadow-xl">
          <div className="flex justify-between items-center mb-6">
            <h2 className="text-lg font-semibold text-slate-200 flex items-center gap-2">
              Live Evaluations
            </h2>
            <button
              onClick={fetchApplications}
              className="p-2 rounded-lg bg-slate-950 hover:bg-slate-800 border border-slate-800 text-slate-400 hover:text-slate-200 transition-colors"
              title="Refresh Data"
            >
              <RefreshCw className={`w-4 h-4 ${loading ? 'animate-spin' : ''}`} />
            </button>
          </div>

          <div className="overflow-x-auto">
            <table className="w-full text-left text-sm text-slate-300">
              <thead className="bg-slate-950 text-slate-400 text-xs uppercase border-b border-slate-800">
                <tr>
                  <th className="p-3">ID</th>
                  <th className="p-3">Applicant</th>
                  <th className="p-3">Income</th>
                  <th className="p-3">Requested</th>
                  <th className="p-3">Status</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-slate-800/50">
                {applications.length === 0 ? (
                  <tr>
                    <td colSpan="5" className="text-center py-8 text-slate-500">
                      No credit applications found. Submit one to see AI in action.
                    </td>
                  </tr>
                ) : (
                  applications.map((app) => (
                    <tr key={app.id} className="hover:bg-slate-800/30 transition-colors">
                      <td className="p-3 font-mono text-xs text-slate-500">#{app.id}</td>
                      <td className="p-3 font-medium text-slate-200">{app.applicant_name}</td>
                      <td className="p-3 font-mono text-emerald-400">${app.monthly_income?.toLocaleString()}</td>
                      <td className="p-3 font-mono text-slate-300">${app.requested_amount?.toLocaleString()}</td>
                      <td className="p-3">
                        {app.status === 'APPROVED' ? (
                          <span className="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-semibold bg-emerald-950/60 text-emerald-400 border border-emerald-800">
                            <CheckCircle2 className="w-3.5 h-3.5" /> Approved
                          </span>
                        ) : app.status === 'REJECTED' ? (
                          <span className="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-semibold bg-rose-950/60 text-rose-400 border border-rose-800">
                            <XCircle className="w-3.5 h-3.5" /> Rejected
                          </span>
                        ) : (
                          <span className="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-semibold bg-amber-950/60 text-amber-400 border border-amber-800">
                            Pending
                          </span>
                        )}
                      </td>
                    </tr>
                  ))
                )}
              </tbody>
            </table>
          </div>
        </section>
      </main>
    </div>
  );
}